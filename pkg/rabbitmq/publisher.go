package pkgRabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/google/wire"
// 	"github.com/pkg/errors"
// 	amqp "github.com/rabbitmq/amqp091-go"
// 	"golang.org/x/exp/slog"
// )

// const (
// 	PUBLISH_MONDATORY = false
// 	PUBLISH_IMMEDIATE = false

// 	EXACHANGE_NAME    = "mainapp-exchange"
// 	BINDING_KEY       = "mainapp-routing-key"
// 	MESSAGE_TYPE_NAME = "mainapp"
// )

// type Publisherer struct {
// 	exchangeName, bindingKey string
// 	messageTypeName          string
// 	amqpChan                 *amqp.Channel
// 	amqpConn                 *amqp.Connection
// }

// var EventPublishererSet = wire.NewSet(NewPublisherer)

// func NewPublisherer(amqpConn *amqp.Connection, exchangeName string, bindingKey string, messageTypeName string) *Publisherer {
// 	ch, err := amqpConn.Channel()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ch.Close()

// 	pub := &Publisherer{
// 		amqpConn:        amqpConn,
// 		amqpChan:        ch,
// 		exchangeName:    exchangeNam
// 		bindingKey:      bindingKey,
// 		messageTypeName: messageTypeName,
// 	}

// 	return pub
// }

// func (p *Publisherer) PublishEvents(ctx context.Context, events []any) error {
// 	for _, e := range events {
// 		b, err := json.Marshal(e)
// 		if err != nil {
// 			return errors.Wrap(err, "publisher-json.Marshal")
// 		}

// 		err = p.Publish(ctx, b, "text/plain")
// 		if err != nil {
// 			return errors.Wrap(err, "publisher-pub.Publish")
// 		}
// 	}
// 	return nil
// }

// // Publish message.
// func (p *Publisherer) Publish(ctx context.Context, body []byte, contentType string) error {
// 	ch, err := p.amqpConn.Channel()
// 	if err != nil {
// 		return errors.Wrap(err, "CreateChannel")
// 	}
// 	defer ch.Close()

// 	slog.Info("publish message", "exchange", p.exchangeName, "routing_key", p.bindingKey)

// 	if err := ch.PublishWithContext(
// 		ctx,
// 		p.exchangeName,
// 		p.bindingKey,
// 		PUBLISH_MONDATORY,
// 		PUBLISH_IMMEDIATE,
// 		amqp.Publishing{
// 			ContentType:  contentType,
// 			DeliveryMode: amqp.Persistent,
// 			MessageId:    uuid.New().String(),
// 			Timestamp:    time.Now(),
// 			Body:         body,
// 			Type:         p.messageTypeName,
// 		},
// 	); err != nil {
// 		return errors.Wrap(err, "ch.Publish")
// 	}

// 	return nil
// }

type Publisher struct {
	conn *amqp091.Connection
}

func NewPublisher(conn *amqp091.Connection) *Publisher {
	return &Publisher{
		conn: conn,
	}
}

func (p *Publisher) MakeChannel() (*amqp091.Channel, error) {
	ch, err := p.conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (p *Publisher) Publish(ctx context.Context, qName string, data []byte, eventType string, cType string) error {
	channel, err := p.MakeChannel()
	if err != nil {
		return err
	}

	channel.Qos(1, 0, false)

	return channel.PublishWithContext(
		ctx,
		"",
		qName,
		false,
		false,
		amqp091.Publishing{
			ContentType: cType,
			Type:        eventType,
			Body:        data,
		},
	)
}
