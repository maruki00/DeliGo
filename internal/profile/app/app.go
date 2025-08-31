package app

import (
	"context"

	grpc_services "github.com/maruki00/deligo/internal/profile/app/service"
	contracts "github.com/maruki00/deligo/internal/profile/domain/contract"
	pkgCqrs "github.com/maruki00/deligo/pkg/cqrs"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	db          *pkgPostgres.PGHandler
	ProfileSVC  *grpc_services.ProfileService
	CommandBus  *pkgCqrs.CommandBus
	QueryBus    *pkgCqrs.QueryBus

	ProfileRepo contracts.IPorofileRepository
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



	profileRepo := repository.NewProfileRepository(db)
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
