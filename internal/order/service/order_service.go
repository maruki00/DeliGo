package service

import (
	"context"
	"errors"
	"time"

	"github.com/maruki00/deligo/internal/order/domain"
	"github.com/maruki00/deligo/internal/order/rabbitmq"
	"github.com/maruki00/deligo/internal/order/requests"

	"github.com/google/uuid"
)

type OrderService struct {
	repo *domain.OrderRepository
	mq   *rabbitmq.RabbitMQClient
}

func NewOrderService(repo *domain.OrderRepository, mq *rabbitmq.RabbitMQClient) *OrderService {
	return &OrderService{repo: repo, mq: mq}
}

func (s *OrderService) CreateOrder(ctx context.Context, req requests.CreateOrderRequest) (*domain.Order, error) {
	var totalAmount float64
	orderItems := make([]domain.OrderItem, len(req.Items))

	orderID := uuid.New().String()

	for i, item := range req.Items {
		totalAmount += item.Price * float64(item.Quantity)
		orderItems[i] = domain.OrderItem{
			OrderID:     orderID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Price:       item.Price,
			Quantity:    item.Quantity,
		}
	}

	order := &domain.Order{
		ID:              orderID,
		CustomerID:      req.CustomerID,
		RestaurantID:    req.RestaurantID,
		Status:          domain.StatusPendingPayment,
		TotalAmount:     totalAmount,
		DeliveryAddress: req.DeliveryAddress,
		Items:           orderItems,
	}

	if err := (*s.repo).Create(ctx, order); err != nil {
		return nil, err
	}

	event := domain.OrderEvent{
		EventName: "order.created",
		OrderID:   order.ID,
		Payload:   *order,
		Timestamp: time.Now(),
	}
	_ = s.mq.Publish(ctx, "order.created", event)

	return order, nil
}

func (s *OrderService) ConfirmOrder(ctx context.Context, id string) error {
	err := (*s.repo).UpdateStatus(ctx, id, domain.StatusPendingPayment, domain.StatusPaid)
	if err != nil {
		return err
	}

	// Broadcast event status change asynchronously
	_ = s.mq.Publish(ctx, "order.paid", map[string]string{"order_id": id, "status": string(domain.StatusPaid)})
	return nil
}

func (s *OrderService) AcceptOrder(ctx context.Context, id string, actor string) error {
	order, err := (*s.repo).FindByID(ctx, id)
	if err != nil {
		return err
	}

	var targetStatus domain.OrderStatus
	var expectedCurrentStatus domain.OrderStatus

	if actor == "restaurant" {
		expectedCurrentStatus = domain.StatusPaid
		targetStatus = domain.StatusPreparing
	} else if actor == "courier" {
		expectedCurrentStatus = domain.StatusPreparing

		// Alternative state path shortcut depending on whether standard flow moves directly
		if order.Status == domain.StatusPreparing {
			targetStatus = domain.StatusReadyForPickup
		} else if order.Status == domain.StatusReadyForPickup {
			targetStatus = domain.StatusPickedUp
		} else {
			return errors.New("order not ready for courier manipulation")
		}
		expectedCurrentStatus = order.Status
	} else {
		return errors.New("invalid lifecycle actor configuration")
	}

	err = (*s.repo).UpdateStatus(ctx, id, expectedCurrentStatus, targetStatus)
	if err != nil {
		return err
	}

	_ = s.mq.Publish(ctx, "order.status_changed", map[string]string{"order_id": id, "status": string(targetStatus)})
	return nil
}

func (s *OrderService) CompleteDelivery(ctx context.Context, id string) error {
	err := (*s.repo).UpdateStatus(ctx, id, domain.StatusPickedUp, domain.StatusDelivered)
	if err != nil {
		return err
	}
	_ = s.mq.Publish(ctx, "order.delivered", map[string]string{"order_id": id, "status": string(domain.StatusDelivered)})
	return nil
}

func (s *OrderService) CancelOrder(ctx context.Context, id string) error {
	order, err := (*s.repo).FindByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status == domain.StatusPickedUp || order.Status == domain.StatusDelivered {
		return errors.New("cannot cancel an order once picked up or delivered")
	}

	err = (*s.repo).UpdateStatus(ctx, id, order.Status, domain.StatusCancelled)
	if err != nil {
		return err
	}

	_ = s.mq.Publish(ctx, "order.cancelled", map[string]string{"order_id": id, "status": string(domain.StatusCancelled)})
	return nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	return (*s.repo).FindByID(ctx, id)
}
