package pkgRabbitmq

import (
	"errors"
	"log/slog"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQDSN string

const (
	_MAXTRIES = 3
	_SLEEP    = 1
)

func NewRabbitMQ(dsn RabbitMQDSN) (*amqp091.Connection, error) {
	var tries int = 0
	for tries < _MAXTRIES {
		slog.Info("trying to connect to RabbitMQ.")
		conn, err := amqp091.Dial(string(dsn))
		if err == nil {
			slog.Info("connected to rabbitmq")
			return conn, nil
		}
		slog.Error("could not reach the rabbitmq server " + err.Error())
		tries++
		time.Sleep(_SLEEP * time.Second)

	}
	return nil, errors.New("could not connect to rabbitmq server")

}
