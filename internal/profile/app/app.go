package app

import (
	"context"

	grpc_services "github.com/maruki00/deligo/internal/profile/app/service"
	profilecfg "github.com/maruki00/deligo/internal/profile/configs"
	contracts "github.com/maruki00/deligo/internal/profile/domain/contract"
	"github.com/maruki00/deligo/internal/profile/infra/repository"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	db          *pkgPostgres.PGHandler
	ProfileSVC  *grpc_services.ProfileServerService
	ProfileRepo contracts.IPorofileRepository
}

func (app *App) GetDB() any {
	return app.db
}

func InitApp(cfg *profilecfg.Config) (*App, func(), error) {
	db, err := pkgPostgres.NewDB(cfg.Postgres.DSN)
	if err != nil {
		return nil, func() {}, err
	}

	profileRepo := repository.NewProfileRepository(db)
	profileSVC := grpc_services.NewProfileService(profileRepo)

	app := &App{
		db:          db,
		ProfileRepo: profileRepo,
		ProfileSVC:  profileSVC,
	}

	return app, func() {}, nil
}

func (a *App) Worker(ctx context.Context) {
	<-ctx.Done()
}
