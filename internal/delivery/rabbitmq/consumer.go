package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/maruki00/deligo/delivery/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderConfirmedEvent struct {
	OrderID             string  `json:"order_id"`
	RestaurantLatitude  float64 `json:"restaurant_latitude"`
	RestaurantLongitude float64 `json:"restaurant_longitude"`
}

type RabbitMQConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	svc     service.DeliveryService
	url     string
}

func NewRabbitMQConsumer(url string, svc service.DeliveryService) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		url: url,
		svc: svc,
	}
}

func (c *RabbitMQConsumer) Start() {
	var err error
	c.conn, err = amqp.Dial(c.url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	q, err := c.channel.QueueDeclare(
		"order.confirmed",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	msgs, err := c.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)
			var event OrderConfirmedEvent
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Printf("Error unmarshaling event payload: %v", err)
				continue
			}

			err := c.svc.HandleOrderConfirmed(event.OrderID, event.RestaurantLatitude, event.RestaurantLongitude)
			if err != nil {
				log.Printf("Error processing dispatch logic: %v", err)
			} else {
				log.Printf("Order %s successfully processed through courier dispatch pipeline", event.OrderID)
			}
		}
	}()

	log.Printf(" [*] Listening to 'order.confirmed' event streaming channel...")
}

func (c *RabbitMQConsumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
