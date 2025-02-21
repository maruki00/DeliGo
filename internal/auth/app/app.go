package app

import (
	"context"
	"delivery/internal/auth/app/services"
	"delivery/internal/auth/domain/contracts"
	"delivery/internal/auth/domain/ports"
	"delivery/internal/auth/infra/repositories"
	"delivery/internal/auth/userGateWay/adapters/presenters"

	"github.com/go-playground/validator"
	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Repo       contracts.IAuthRepository
	Validate   *validator.Validate
	InputPort  ports.AuthInputPort
	OutputPort ports.AuthOutputPort
}

func InitApp() (*App, error) {

	repo := repositories.NewAuthRepository(nil)
	outPort := &presenters.JsonAuthPresenter{}
	inPort := &services.AuthService{}
	return &App{
		Repo:       repo,
		InputPort:  inPort,
		OutputPort: outPort,
		Validate:   validator.New(),
	}, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

}
