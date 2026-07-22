package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const ExchangeName = "orders_exchange"

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQClient(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQClient{conn: conn, channel: ch}, nil
}

func (r *RabbitMQClient) Publish(ctx context.Context, routingKey string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return r.channel.PublishWithContext(ctx,
		ExchangeName,
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

func (r *RabbitMQClient) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
