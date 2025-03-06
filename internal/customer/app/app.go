package app

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
}

// cfg *configs.Config
func InitApp() (*App, error) {

	return nil, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

}
