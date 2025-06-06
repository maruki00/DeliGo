package app

import (
	"context"
	"deligo/internal/order/domain/contracts"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Repository contracts.IOrderRepository
}

func NewApp() *App {
	return &App{}
}

func (app *App) Worder(ctx context.Context, deligo <-chan amqp091.Delivery) {

}
