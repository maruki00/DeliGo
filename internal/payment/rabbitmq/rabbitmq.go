package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/maruki00/deligo/internal/payment/config"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPCliet struct {
	Conn       *amqp091.Connection
	Ch         *amqp091.Channel
	exchange   string
	routingKey string
}

func NewAMQPClient(cfg *config.Config) (*AMQPCliet, error) {
	conn, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		return nil, fmt.Errorf("Could not establish connection to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to open communication channel with RabbitMQ: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		cfg.RabbitMQ.ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ exchange system: %v", err)
	}

	return &AMQPCliet{
		Conn:       conn,
		Ch:         ch,
		exchange:   cfg.RabbitMQ.ExchangeName,
		routingKey: cfg.RabbitMQ.RoutingKey,
	}, nil
}

func (amqp *AMQPCliet) Consumer() {

}

func (amqp *AMQPCliet) Publish(ctx context.Context, body []byte) error {
	return amqp.Ch.PublishWithContext(
		context.Background(),
		amqp.exchange,
		amqp.routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

}
