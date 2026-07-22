package restaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/restaurant/config"
	"github.com/maruki00/deligo/internal/restaurant/handlers"
	"github.com/maruki00/deligo/internal/restaurant/repositories"
	"github.com/maruki00/deligo/internal/restaurant/routers"
	"github.com/maruki00/deligo/internal/restaurant/services"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	DB         *pkgPostgres.PGHandler
	Repo       repositories.CatalogRepository
	Svc        services.CatalogService
	Handler    *handlers.CatalogHandler
	HTTPServer *gin.Engine
}

func Iinit(cfg *config.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.DB.Dsn)
	if err != nil {
		return nil, func() {
			if db != nil {
				db.Close()
			}
		}, err
	}
	repo := repositories.NewCatalogRepository(db.DB)
	svc := services.NewCatalogService(repo)
	handler := handlers.NewCatalogHandler(svc)
	httpSVC := routers.SetupRouter(handler)

	return &App{
			DB:         db,
			Repo:       repo,
			Svc:        svc,
			Handler:    handler,
			HTTPServer: httpSVC,
		}, func() {
			db.Close()
		}, nil
}
