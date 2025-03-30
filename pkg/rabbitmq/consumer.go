package pkgRabbitmq

import (
	"context"
	"delivery/cmd/auth/configs"

	"github.com/rabbitmq/amqp091-go"
)

// import (
// 	"context"
// 	"log/slog"

// 	"github.com/pkg/errors"
// 	"github.com/rabbitmq/amqp091-go"
// )

// const (

// 	//Consume
// 	CONSUME_AutoAck   = false
// 	CONSUME_Exclusive = false
// 	CONSUME_NoLocal   = false
// 	CONSUME_NoWait    = false

// 	//ExchangeDeclare
// 	EXCHANGE_KINF        = ""
// 	EXCHANGE_DURABLE     = false
// 	EXCHANGE_AUTO_DELETE = false
// 	EXCHANGE_INTERNAL    = false
// 	EXCHANGE_NO_WAIT     = false

// 	//Queue Declare
// 	QUEUE_Durable    = false
// 	QUEUE_AutoDelete = false
// 	QUEUE_Exclusive  = false
// 	QUEUE_NoWait     = false

// 	POOLSIZE = 32

// 	//Qos
// 	PREFETCH_Count  = 5
// 	PREFETCH_Size   = 0
// 	PREFETCH_Global = false
// )

// type Worker func(context context.Context, delivery <-chan amqp091.Delivery)
// type Consumer struct {
// 	amqpConn    *amqp091.Connection
// 	QueueName   string
// 	ExangeName  string
// 	bindingkey  string
// 	ConsumerTag string
// 	PoolSize    int
// }

// func NewConsumer(conn *amqp091.Connection, _bindingkey string, _QueueName string, _ConsumerTag string) *Consumer {
// 	return &Consumer{
// 		amqpConn:    conn,
// 		QueueName:   _QueueName,
// 		bindingkey:  _bindingkey,
// 		ConsumerTag: _ConsumerTag,
// 		PoolSize:    POOLSIZE,
// 	}
// }

// func (c *Consumer) Start(fn Worker) error {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	ch, err := c.createChannel()
// 	if err != nil {
// 		return errors.Wrap(err, "CreateChannel")
// 	}
// 	defer ch.Close()

// 	deliveries, err := ch.Consume(
// 		c.QueueName,
// 		c.ConsumerTag,
// 		CONSUME_AutoAck,
// 		CONSUME_Exclusive,
// 		CONSUME_NoLocal,
// 		CONSUME_NoWait,
// 		nil,
// 	)
// 	if err != nil {
// 		return errors.Wrap(err, "Consume")
// 	}

// 	forever := make(chan bool)

// 	for i := 0; i < c.PoolSize; i++ {
// 		go fn(ctx, deliveries)
// 	}

// 	chanErr := <-ch.NotifyClose(make(chan *amqp091.Error))
// 	slog.Error("ch.NotifyClose", chanErr)
// 	<-forever

// 	return chanErr
// }

// // CreateChannel Consume messages.
// func (c *Consumer) createChannel() (*amqp091.Channel, error) {
// 	ch, err := c.amqpConn.Channel()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error amqpConn.Channel")
// 	}

// 	slog.Info("declaring exchange", "exchange_name", c.ExangeName)
// 	err = ch.ExchangeDeclare(
// 		c.ExangeName,
// 		EXCHANGE_KINF,
// 		EXCHANGE_DURABLE,
// 		EXCHANGE_AUTO_DELETE,
// 		EXCHANGE_INTERNAL,
// 		EXCHANGE_NO_WAIT,
// 		nil,
// 	)

// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
// 	}

// 	queue, err := ch.QueueDeclare(
// 		c.QueueName,
// 		QUEUE_Durable,
// 		QUEUE_AutoDelete,
// 		QUEUE_Exclusive,
// 		QUEUE_NoWait,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error ch.QueueDeclare")
// 	}

// 	err = ch.QueueBind(
// 		queue.Name,
// 		c.bindingkey,
// 		c.ExangeName,
// 		QUEUE_NoWait,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error ch.QueueBind")
// 	}

// 	slog.Info("queue bound to exchange, starting to consume from queue", "consumer_tag", c.ConsumerTag)

// 	err = ch.Qos(
// 		PREFETCH_Count,  // prefetch count
// 		PREFETCH_Size,   // prefetch size
// 		PREFETCH_Global, // global
// 	)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error ch.Qos")
// 	}

// 	return ch, nil
// }

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
	cfg *configs.Config,
) *Consumer {
	return &Consumer{
		conn:       conn,
		queuesName: cfg.Queues,
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

func (c *Consumer) StartConsumer(worker func(ctx context.Context, delivery <-chan amqp091.Delivery)) error {

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

		delivery, err := channel.Consume(
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
			go worker(ctx, delivery)
		}
	}
	forever := make(chan bool)
	amqpError := channel.NotifyClose(make(chan *amqp091.Error))
	<-forever
	return <-amqpError
}
