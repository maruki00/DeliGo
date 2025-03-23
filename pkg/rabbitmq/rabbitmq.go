package pkgRabbitmq

import (
	"errors"
	"fmt"
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
	var cnx *amqp091.Connection
	var tries int = 0
	for tries < _MAXTRIES {
		slog.Info("trying to connect to RabbitMQ.")
		fmt.Println("dsn : ", string(dsn))
		conn, err := amqp091.Dial(string(dsn))
		if err != nil {
			slog.Error("could not reach the rabbitmq server " + err.Error())
			tries++
			time.Sleep(_SLEEP * time.Second)
			continue
		}
		cnx = conn
		break
	}
	if cnx == nil {
		slog.Error("could not connect the rabbitmq server")
		return nil, errors.New("could not connect to rabbitmq server")
	}

	slog.Info("connected to rabbitmq")
	return cnx, nil

}
