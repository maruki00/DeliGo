package pkgRabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

const (
	POOL_SIZE = 32
)

type Consumer struct {
	conn       *amqp091.Connection
	queuesName []string
	poolSize   int
}

func NewConsumer(
	conn *amqp091.Connection,
	queuesName []string,
) *Consumer {
	return &Consumer{
		conn:       conn,
		queuesName: queuesName,
		poolSize:   POOL_SIZE,
	}
}

func (c *Consumer) MakeChannel() (*amqp091.Channel, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *Consumer) StartConsumer(worker func(ctx context.Context, github.com/maruki00/deligo <-chan amqp091.Delivery)) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	channel, err := c.MakeChannel()
	if err != nil {
		return err
	}

	channel.Qos(1, 0, false)
	for _, queueName := range c.queuesName {
		queue, err := channel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		github.com/maruki00/deligo, err := channel.Consume(
			queue.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		for i := 0; i < int(c.poolSize/len(c.queuesName)); i++ {
			go worker(ctx, github.com/maruki00/deligo)
		}
	}
	forever := make(chan bool)
	amqpError := channel.NotifyClose(make(chan *amqp091.Error))
	<-forever
	return <-amqpError
}
