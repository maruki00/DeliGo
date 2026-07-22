package messaging

import (
	"encoding/json"
	"log"

	"github.com/maruki00/deligo/internal/analytic/config"
	"github.com/maruki00/deligo/internal/analytic/domain"
	"github.com/maruki00/deligo/internal/analytic/repository"

	"github.com/rabbitmq/amqp091-go"
)

type EventConsumer struct {
	channel      *amqp091.Channel
	cfg          *config.Config
	analyticRepo repository.AnalyticRepository
}

func NewEventConsumer(ch *amqp091.Channel, cfg *config.Config, repo repository.AnalyticRepository) *EventConsumer {
	return &EventConsumer{channel: ch, cfg: cfg, analyticRepo: repo}
}

func (c *EventConsumer) StartConsuming() {
	msgs, err := c.channel.Consume(
		c.cfg.QueueName,
		"feedback-analytics-worker", // Consumer reference tag
		false,                       // Auto-ack set to false to guarantee processing durability
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Critical: Background worker failed to consume queue stream: %v", err)
	}

	go func() {
		log.Println("Asynchronous Analytics Worker running. Awaiting payloads...")
		for d := range msgs {
			var feedback domain.Feedback
			if err := json.Unmarshal(d.Body, &feedback); err != nil {
				log.Printf("Error decoding event body payload: %v", err)
				d.Nack(false, false) // Reject and discard dead lettering drop
				continue
			}

			log.Printf("[Worker] Event picked up for processing. Aggregating Product ID: %s", feedback.ProductID)

			// Process aggregation updating projection views
			if err := c.analyticRepo.UpdateMetrics(feedback.ProductID); err != nil {
				log.Printf("Worker error scaling updates calculations: %v", err)
				d.Nack(false, true) // Requeue tracking error failure state
				continue
			}

			d.Ack(false) // Safe execution context acknowledgment
		}
	}()
}
