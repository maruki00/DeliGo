package courier

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/courier/config"
	"github.com/maruki00/deligo/internal/courier/domain"
	"github.com/maruki00/deligo/internal/courier/handlers"
	"github.com/maruki00/deligo/internal/courier/rabbitmq"
	"github.com/maruki00/deligo/internal/courier/repository"
	"github.com/maruki00/deligo/internal/courier/router"
	"github.com/maruki00/deligo/internal/courier/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	Repo       domain.CourierRepository
	Svc        domain.CourierService
	Handler    *handlers.CourierHandler
	HTTPServer *gin.Engine
	Consumer   *rabbitmq.RabbitMQConsumer
}

func Init(cfg *config.Config) (*App, func(), error) {
	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {}, err
	}

	repo := repository.NewCourierRepository(db.DB)
	svc := service.NewCourierService(repo)
	handler := handlers.NewCourierHandler(svc)
	consumer := rabbitmq.NewRabbitMQConsumer(cfg.RabbitMQ.URL, svc)

	return &App{
			Repo:       repo,
			Svc:        svc,
			Handler:    handler,
			Consumer:   consumer,
			HTTPServer: router.SetupRouter(handler),
		}, func() {
			db.Close()
		}, nil
}
