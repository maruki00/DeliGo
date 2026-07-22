package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	mu     sync.RWMutex
	url    string
	exName string
}

func NewPublisher(url string) (*Publisher, error) {
	p := &Publisher{
		url:    url,
		exName: "user.events",
	}
	if err := p.connect(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Publisher) connect() error {
	var err error
	p.conn, err = amqp.Dial(p.url)
	if err != nil {
		return fmt.Errorf("rabbitmq dial failure: %w", err)
	}

	p.ch, err = p.conn.Channel()
	if err != nil {
		p.conn.Close()
		return fmt.Errorf("failed to open rabbitmq channel: %w", err)
	}

	return p.ch.ExchangeDeclare(
		p.exName,
		"topic",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,
	)
}

func (p *Publisher) PublishEvent(ctx context.Context, routingKey string, payload interface{}) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return p.ch.PublishWithContext(c,
		p.exName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         body,
		},
	)
}

func (p *Publisher) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ch != nil {
		p.ch.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
