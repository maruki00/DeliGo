package app

import (
	"context"
	"deligo/cmd/auth/configs"
	"deligo/internal/auth/app/services"
	"deligo/internal/auth/domain/contracts"
	"deligo/internal/auth/domain/ports"
	"deligo/internal/auth/infra/repositories"
	"deligo/internal/auth/userGateWay/adapters/presenters"

	"github.com/go-playground/validator"
	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	Repo       contracts.IAuthRepository
	Validate   *validator.Validate
	InputPort  ports.AuthInputPort
	OutputPort ports.AuthOutputPort
}

func InitApp(cfg *configs.Config) (*App, error) {

	repo, err := repositories.NewAuthRepository(cfg.Postgres.Dsn)
	if err != nil {
		return nil, err
	}
	outPort := &presenters.JsonAuthPresenter{}
	inPort := &services.AuthService{}
	return &App{
		Repo:       repo,
		InputPort:  inPort,
		OutputPort: outPort,
		Validate:   validator.New(),
	}, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.deligo) {

}
