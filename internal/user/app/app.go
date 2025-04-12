package app

import (
	"context"
	"delivery/cmd/user/configs"
	grpc_services "delivery/internal/user/app/grpc/services"
	"delivery/internal/user/infra/repositories"
	pkgPostgres "delivery/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db          *pkgPostgres.PGHandler
	UserRepo    *repositories.UserRepository
	ProfileRepo *repositories.ProfileRepository
	UserSVC     *grpc_services.UserService
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *configs.Config) (*App, func(), error) {

	fmt.Println("dsn : ", cfg.Postgres.Dsn)
	db, err := pkgPostgres.NewPG(cfg.Postgres.Dsn)
	if err != nil {
		return nil, func() {}, err
	}
	userRepo := repositories.NewUserRepository(*db)
	userSVC := grpc_services.NewUserService(userRepo)
	app := &App{
		db:          db,
		UserRepo:    userRepo,
		ProfileRepo: nil,
		UserSVC:     userSVC,
	}
	return app, func() { _ = db.DB.Close() }, nil
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
