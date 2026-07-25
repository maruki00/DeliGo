package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/maruki00/deligo/internal/payment/domain"
	"github.com/maruki00/deligo/internal/payment/rabbitmq"
	"github.com/maruki00/deligo/internal/payment/requests"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

type PaymentSucceededEvent struct {
	PaymentID  string  `json:"payment_id"`
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
}

type PaymentService struct {
	repo       domain.PaymentRepository
	amqpClient *rabbitmq.AMQPCliet
}

func NewPaymentService(repo domain.PaymentRepository, amqp *rabbitmq.AMQPCliet) *PaymentService {
	return &PaymentService{
		repo:       repo,
		amqpClient: amqp,
	}
}

func (s *PaymentService) ProcessCharge(ctx context.Context, req *requests.ChargeRequest) (*domain.Payment, error) {

	p := &domain.Payment{
		OrderID:    req.OrderID,
		CustomerID: req.CustomerID,
		Amount:     req.Amount,
		Currency:   req.Currency,
		Status:     domain.StatusInitiated,
	}

	if err := s.repo.Create(p); err != nil {
		return nil, fmt.Errorf("failed to save initiated log: %w", err)
	}

	stripeAmount := int64(math.Round(req.Amount * 100))

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(stripeAmount),
		Currency: stripe.String(req.Currency),
		Source:   &stripe.SourceParams{Token: stripe.String(req.Token)},
		// MetaData: map[string]string{
		// 	"payment_id":  p.ID,
		// 	"order_id":    p.OrderID,
		// 	"customer_id": p.CustomerID,
		// },
	}

	stripeCharge, err := charge.New(chargeParams)
	if err != nil {
		log.Printf("Stripe transaction failed for payment %s: %v", p.ID, err)
		_ = s.repo.UpdateStatus(p.ID, domain.StatusFailed, nil)
		p.Status = domain.StatusFailed
		return p, fmt.Errorf("stripe processing failure: %w", err)
	}

	err = s.repo.UpdateStatus(p.ID, domain.StatusSucceeded, &stripeCharge.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to finalize transaction record: %w", err)
	}
	p.Status = domain.StatusSucceeded
	p.StripeChargeID = &stripeCharge.ID

	s.publishSucceededEvent(p)

	return p, nil
}

func (s *PaymentService) ProcessRefund(ctx context.Context, req *requests.RefundRequest) (*domain.Payment, error) {

	return nil, nil
}

func (s *PaymentService) HandleWebhook(ctx context.Context, payload []byte, sigHeader string, webhookSecret string) error {
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
		if err == nil && p.Status == domain.StatusInitiated {
			_ = s.repo.UpdateStatus(p.ID, domain.StatusSucceeded, &c.ID)
			p.Status = domain.StatusSucceeded
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
		_ = s.repo.UpdateStatus(paymentID, domain.StatusRefunded, nil)
	}
	return nil
}

func (s *PaymentService) publishSucceededEvent(p *domain.Payment) {
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

	err = s.amqpClient.Publish(context.TODO(), body)
	if err != nil {
		log.Printf("Failed to dispatch payment.succeeded message to RabbitMQ: %v", err)
	} else {
		log.Printf("Successfully published payment.succeeded for order: %s", p.OrderID)
	}
}
