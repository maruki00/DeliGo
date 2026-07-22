package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/maruki00/deligo/internal/notifier/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	connStr string
	svc     service.Service
	conn    *amqp.Connection
	channel *amqp.Channel
}

type NotificationEvent struct {
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// NewRabbitMQConnection handles structural instantiation of your event infrastructure handlers
func NewRabbitMQConnection(connStr string, svc service.Service) *RabbitMQConnection {
	return &RabbitMQConnection{
		connStr: connStr,
		svc:     svc,
	}
}

// StartConsuming initializes an asynchronous concurrent listener loop looking for event packets
func (rc *RabbitMQConnection) StartConsuming(queueName string) {
	go func() {
		for {
			log.Println("Attempting to connect to RabbitMQ broker infrastructure...")
			conn, err := amqp.Dial(rc.connStr)
			if err != nil {
				log.Printf("Failed to establish raw stream line to messaging engine: %v. Retrying in 5s...", err)
				time.Sleep(5 * time.Second)
				continue
			}
			rc.conn = conn

			ch, err := conn.Channel()
			if err != nil {
				log.Printf("Failed to unlock dynamic handling session channel: %v", err)
				conn.Close()
				time.Sleep(5 * time.Second)
				continue
			}
			rc.channel = ch

			q, err := ch.QueueDeclare(
				queueName,
				true,  // Durable
				false, // Delete when unused
				false, // Exclusive
				false, // No-wait
				nil,   // Arguments
			)
			if err != nil {
				log.Printf("Failed to assert consumer target queue layout criteria: %v", err)
				ch.Close()
				conn.Close()
				time.Sleep(5 * time.Second)
				continue
			}

			msgs, err := ch.Consume(
				q.Name,
				"",    // Consumer tag identifier
				true,  // Auto-Ack
				false, // Exclusive
				false, // No-local
				false, // No-wait
				nil,   // Args
			)
			if err != nil {
				log.Printf("Queue operational processing flow interface execution failure: %v", err)
				ch.Close()
				conn.Close()
				time.Sleep(5 * time.Second)
				continue
			}

			log.Printf("RabbitMQ Consumer connected. Listening on queue: %s", queueName)

			// Loop breaks naturally if the channel closes out underneath us due to network errors
			for d := range msgs {
				var event NotificationEvent
				if err := json.Unmarshal(d.Body, &event); err != nil {
					log.Printf("Failed to unmarshal notification incoming wire block payload: %v", err)
					continue
				}

				log.Printf("[EVENT RECEIVED] Processing order status payload for user: %s", event.UserID)
				if err := rc.svc.ProcessInboundNotification(event.UserID, event.Title, event.Body); err != nil {
					log.Printf("System background engine failed logic distribution mapping execution steps: %v", err)
				}
			}

			log.Println("RabbitMQ stream connection broken or terminated. Re-initiating cycle loop...")
			ch.Close()
			conn.Close()
		}
	}()
}

// PublishNotification facilitates outbound events to specific queues if needed inside the application
func (rc *RabbitMQConnection) PublishNotification(ctx context.Context, queueName string, event NotificationEvent) error {
	if rc.channel == nil {
		log.Println("Cannot publish message: RabbitMQ channel connection mapping is uninitialized")
		return amqp.ErrClosed
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Double-checks that the target queue topology exists prior to writing blocks
	_, err = rc.channel.QueueDeclare(
		queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	return rc.channel.PublishWithContext(
		ctx,
		"",        // Exchange
		queueName, // Routing key matches queue directly
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent, // Restructure payload states for absolute storage validation
			Body:         body,
		},
	)
}

// Close gracefully shuts down existing background channels out of system runtime memory maps
func (rc *RabbitMQConnection) Close() {
	if rc.channel != nil {
		rc.channel.Close()
	}
	if rc.conn != nil {
		rc.conn.Close()
	}
}
