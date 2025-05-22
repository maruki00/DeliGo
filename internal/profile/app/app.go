package app

import (
	"context"
	"deligo/cmd/user/configs"
	"deligo/internal/profile/app/profile/commands"
	"deligo/internal/profile/app/profile/handlers"
	"deligo/internal/profile/app/profile/queries"
	grpc_services "deligo/internal/profile/app/usecases"
	"deligo/internal/profile/domain/contracts"
	"deligo/internal/profile/infra/repositories"

	pkgCqrs "deligo/pkg/cqrs"
	pkgPostgres "deligo/pkg/postgres"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

type App struct {
	db          *pkgPostgres.PGHandler
	ProfileRepo contracts.IPorofileRepository
	ProfileSVC  *grpc_services.ProfileService
	CommandBus  *pkgCqrs.CommandBus
	QueryBus    *pkgCqrs.QueryBus
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

	commandBus := pkgCqrs.NewCommandBus()
	queryBus := pkgCqrs.NewQueryBus()

	commandBus.Register(&commands.DiscableProfileCommand{}, handlers.NewDisableProfileHandler(profileRepo))
	commandBus.Register(&commands.SaveProfileCommand{}, handlers.NewSaveProfileHandler(profileRepo))
	commandBus.Register(&commands.UpdateProfileAvatarCommand{}, handlers.NewUpdateProfileAvatarHandler(profileRepo))
	commandBus.Register(&commands.DiscableProfileCommand{}, handlers.NewDisableProfileHandler(profileRepo))
	queryBus.Register(&queries.GetOneProfileQuery{}, handlers.NewGetOneProfileHandler(profileRepo))
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
