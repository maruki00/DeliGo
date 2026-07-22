package restaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/cmd/restaurant/config"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
	"github.com/maruki00/deligo/restaurant/handlers"
	"github.com/maruki00/deligo/restaurant/repositories"
	"github.com/maruki00/deligo/restaurant/routers"
	"github.com/maruki00/deligo/restaurant/services"
)

type App struct {
	DB         *pkgPostgres.PGHandler
	Repo       repositories.CatalogRepository
	Svc        services.CatalogService
	Handler    *handlers.CatalogHandler
	HTTPServer *gin.Engine
}

func Iinit(cfg *config.Cfg) (*App, func(), error) {

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
