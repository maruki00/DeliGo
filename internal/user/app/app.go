package app

import (
	"context"
	"delivery/cmd/user/configs"
	"delivery/internal/user/domain/contracts"
	pkgPostgres "delivery/pkg/postgres"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	cgf       configs.Config
	DBHandler pkgPostgres.PGHandler
	Repo      contracts.IAuthRepository
}

func InitApp(cfg *configs.Config) (*App, error) {

	return nil, nil
}

func (a *App) Worker(ctx context.Context, deivery <-chan amqp091.Delivery) {

	for {
		select {
		case <-ctx.Done():
			slog.Info("Shuting Down the client.")
			break
		default:
			slog.Info("default interception ....")
		}
	}
	// forever := struct{}{}
	// <-forever
}
