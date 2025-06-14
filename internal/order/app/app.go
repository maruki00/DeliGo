package app

import (
	"context"
	"github.com/maruki00/deligo/internal/order/domain/contracts"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Repository contracts.IOrderRepository
}

func NewApp() *App {
	return &App{}
}

func (app *App) Worder(ctx context.Context, github.com/maruki00/deligo <-chan amqp091.Delivery) {

}
