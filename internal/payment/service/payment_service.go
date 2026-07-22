package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/maruki00/deligo/internal/payment/models"
	"github.com/maruki00/deligo/internal/payment/requests"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/refund"
)

type PaymentSucceededEvent struct {
	PaymentID  string  `json:"payment_id"`
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
}

type PaymentService interface {
	ProcessCharge(ctx context.Context, req *requests.ChargeRequest) (*models.Payment, error)
	ProcessRefund(ctx context.Context, req *requests.RefundRequest) (*models.Payment, error)
	HandleWebhook(ctx context.Context, payload []byte, sigHeader string, webhookSecret string) error
}

type paymentService struct {
	repo       repo.PaymentRepository
	ch         *amqp.Channel
	exchange   string
	routingKey string
}

func NewPaymentService(repo repo.PaymentRepository, ch *amqp.Channel, exchange string, routingKey string) PaymentService {
	return &paymentService{
		repo:       repo,
		ch:         ch,
		exchange:   exchange,
		routingKey: routingKey,
	}
}

func (s *paymentService) ProcessCharge(ctx context.Context, req *requests.ChargeRequest) (*models.Payment, error) {
	p := &models.Payment{
		OrderID:    req.OrderID,
		CustomerID: req.CustomerID,
		Amount:     req.Amount,
		Currency:   req.Currency,
		Status:     models.StatusInitiated,
	}

	if err := s.repo.Create(p); err != nil {
		return nil, fmt.Errorf("failed to save initiated log: %w", err)
	}

	stripeAmount := int64(math.Round(req.Amount * 100))

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(stripeAmount),
		Currency: stripe.String(req.Currency),
		Source:   &stripe.SourceParams{Token: stripe.String(req.Token)},
		Metadata: map[string]string{
			"payment_id":  p.ID,
			"order_id":    p.OrderID,
			"customer_id": p.CustomerID,
		},
	}

	stripeCharge, err := charge.New(chargeParams)
	if err != nil {
		log.Printf("Stripe transaction failed for payment %s: %v", p.ID, err)
		_ = s.repo.UpdateStatus(p.ID, models.StatusFailed, nil)
		p.Status = models.StatusFailed
		return p, fmt.Errorf("stripe processing failure: %w", err)
	}

	err = s.repo.UpdateStatus(p.ID, models.StatusSucceeded, &stripeCharge.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to finalize transaction record: %w", err)
	}
	p.Status = models.StatusSucceeded
	p.StripeChargeID = &stripeCharge.ID

	s.publishSucceededEvent(p)

	return p, nil
}

func (s *paymentService) ProcessRefund(ctx context.Context, req *requests.RefundRequest) (*models.Payment, error) {
	p, err := s.repo.FindByOrderID(req.OrderID)
	if err != nil {
		return nil, fmt.Errorf("transaction log not found for order: %w", err)
	}

	if p.Status != models.StatusSucceeded || p.StripeChargeID == nil {
		return nil, errors.New("cannot refund transaction unless status is succeeded")
	}

	refundParams := &stripe.RefundParams{
		Charge: stripe.String(*p.StripeChargeID),
	}

	_, err = refund.New(refundParams)
	if err != nil {
		return nil, fmt.Errorf("stripe refund rejected: %w", err)
	}

	err = s.repo.UpdateStatus(p.ID, models.StatusRefunded, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to mark status as refunded locally: %w", err)
	}
	p.Status = models.StatusRefunded

	return p, nil
}

func (s *paymentService) HandleWebhook(ctx context.Context, payload []byte, sigHeader string, webhookSecret string) error {
	var event stripe.Event
	err := json.Unmarshal(payload, &event)
	if err != nil {
		return fmt.Errorf("invalid webhook json data payload: %w", err)
	}

	switch event.Type {
	case "charge.succeeded":
		var c stripe.Charge
		err := json.Unmarshal(event.Data.Raw, &c)
		if err != nil {
			return fmt.Errorf("failed to parse stripe charge metadata: %w", err)
		}
		paymentID, ok := c.Metadata["payment_id"]
		if !ok {
			return nil
		}
		p, err := s.repo.FindByID(paymentID)
		if err == nil && p.Status == models.StatusInitiated {
			_ = s.repo.UpdateStatus(p.ID, models.StatusSucceeded, &c.ID)
			p.Status = models.StatusSucceeded
			p.StripeChargeID = &c.ID
			s.publishSucceededEvent(p)
		}
	case "charge.refunded":
		var c stripe.Charge
		err := json.Unmarshal(event.Data.Raw, &c)
		if err != nil {
			return fmt.Errorf("failed to parse stripe refund metadata: %w", err)
		}
		paymentID, ok := c.Metadata["payment_id"]
		if !ok {
			return nil
		}
		_ = s.repo.UpdateStatus(paymentID, models.StatusRefunded, nil)
	}

	return nil
}

func (s *paymentService) publishSucceededEvent(p *models.Payment) {
	evt := PaymentSucceededEvent{
		PaymentID:  p.ID,
		OrderID:    p.OrderID,
		CustomerID: p.CustomerID,
		Amount:     p.Amount,
		Currency:   p.Currency,
	}

	body, err := json.Marshal(evt)
	if err != nil {
		log.Printf("Failed to serialize RabbitMQ message payload: %v", err)
		return
	}

	err = s.ch.PublishWithContext(
		context.Background(),
		s.exchange,
		s.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to dispatch payment.succeeded message to RabbitMQ: %v", err)
	} else {
		log.Printf("Successfully published payment.succeeded for order: %s", p.OrderID)
	}
}
