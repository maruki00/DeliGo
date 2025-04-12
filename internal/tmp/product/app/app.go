package app

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
}

func Init() *App {
	return &App{}
}
func (app *App) Worder(ctx context.Context, delivery <-chan amqp091.Delivery) {

}
