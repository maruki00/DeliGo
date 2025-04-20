package app

import (
	"context"
	"deligo/cmd/profile/configs"
	grpc_services "deligo/internal/profile/app/grpc/services"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/infra/repositories"

	pkgPostgres "deligo/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db          *pkgPostgres.PGHandler
	ProfileRepo contracts.IPorofileRepository
	ProfileSVC  *grpc_services.ProfileService
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *configs.Config) (*App, func(), error) {

	fmt.Println("dsn : ", cfg.Postgres.Dsn)
	db, err := pkgPostgres.NewDB(cfg.Postgres.Dsn)
	if err != nil {
		return nil, func() {}, err
	}

	profileRepo := repositories.NewProfileRepository(db)
	profileSVC := grpc_services.NewProfileService(profileRepo)

	app := &App{
		db:          db,
		ProfileRepo: profileRepo,
		ProfileSVC:  profileSVC,
	}

	return app, func() {}, nil
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
