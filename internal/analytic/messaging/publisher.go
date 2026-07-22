package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/maruki00/deligo/internal/analytic/config"
	"github.com/maruki00/deligo/internal/analytic/domain"

	"github.com/rabbitmq/amqp091-go"
)

type EventPublisher interface {
	PublishFeedbackCreated(feedback *domain.Feedback) error
}

type eventPublisher struct {
	channel *amqp091.Channel
	cfg     *config.Config
}

func NewEventPublisher(ch *amqp091.Channel, cfg *config.Config) EventPublisher {
	return &eventPublisher{channel: ch, cfg: cfg}
}

func (p *eventPublisher) PublishFeedbackCreated(feedback *domain.Feedback) error {
	body, err := json.Marshal(feedback)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.channel.PublishWithContext(ctx,
		p.cfg.ExchangeName,
		p.cfg.RoutingKey,
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent,
			Body:         body,
		},
	)
}
