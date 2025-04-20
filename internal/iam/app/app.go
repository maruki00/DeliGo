package app

import (
	"context"
	"deligo/cmd/user/configs"
	grpc_services "deligo/internal/iam/app/grpc/services"
	"deligo/internal/iam/domain/contracts"
	"deligo/internal/iam/infra/repositories"
	pkgPostgres "deligo/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db             *pkgPostgres.PGHandler
	UserRepo       contracts.IUserRepository
	GroupRepo      contracts.IGroupRepository
	PermissionRepo contracts.IPermissionRepository
	PolicyRepo     contracts.IPolicyRepository

	UserSVC *grpc_services.UserService
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

	userRepo := repositories.NewUserRepository(*db)
	userSVC := grpc_services.NewUserService(userRepo)

	profileRepo := repositories.NewProfileRepository(*db)
	profileSVC := grpc_services.NewProfileService(profileRepo)

	app := &App{
		db:          db,
		UserRepo:    userRepo,
		ProfileRepo: profileRepo,
		UserSVC:     userSVC,
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
