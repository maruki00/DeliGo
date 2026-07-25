package messaging

import (
	"log"
	"time"

	"github.com/maruki00/deligo/internal/analytic/config"

	"github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(cfg *config.Config) (*amqp091.Connection, *amqp091.Channel) {
	var conn *amqp091.Connection
	var err error

	for i := 0; i < 5; i++ {
		conn, err = amqp091.Dial(cfg.RabbitMQURL)
		if err == nil {
			break
		}
		log.Printf("RabbitMQ connection failure. Retrying in 3 seconds... (%d/5)", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Fatal: Could not connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Fatal: Failed to open a channel over broker: %v", err)
	}

	// Declare Exchange
	err = ch.ExchangeDeclare(
		cfg.ExchangeName,
		"topic",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	// Declare Target Worker Queue
	q, err := ch.QueueDeclare(
		cfg.QueueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// Bind Queue to Topic Exchange Topology
	err = ch.QueueBind(
		q.Name,
		cfg.RoutingKey,
		cfg.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind topology paths: %v", err)
	}

	log.Println("RabbitMQ Topology configured successfully.")
	return conn, ch
}
