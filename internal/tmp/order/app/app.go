package app

import (
	"context"
	"deligo/internal/order/app/services"
	"deligo/internal/order/domain/contracts"
	"deligo/internal/order/domain/ports"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Service    *services.OrderService
	Repository contracts.IOrderRepository
	InputPort  ports.OrderInputPort
	OutputPort ports.OrderOutputPort
}

func NewApp(
	Service *services.OrderService,
	Repository contracts.IOrderRepository,
	InputPort ports.OrderInputPort,
	OutputPort ports.OrderOutputPort,
) *App {
	return &App{
		Service:    Service,
		Repository: Repository,
		InputPort:  InputPort,
		OutputPort: OutputPort,
	}
}

func (app *App) Worder(ctx context.Context, deligo <-chan amqp091.deligo) {

}
