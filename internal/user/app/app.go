package app

import (
	"context"
	"delivery/cmd/user/configs"
	"delivery/internal/user/app/services"
	"delivery/internal/user/domain/contracts"
	"delivery/internal/user/infra/repositories"
	"delivery/internal/user/userGateWay/adapters/presenters"

	"github.com/go-playground/validator"
	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	cgf  configs.Config
	db   pkgPostgres.PGHandler
	Repo contracts.IAuthRepository
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

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

}
