package order

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/order/config"
	"github.com/maruki00/deligo/internal/order/domain"
	"github.com/maruki00/deligo/internal/order/handlers"
	"github.com/maruki00/deligo/internal/order/rabbitmq"
	"github.com/maruki00/deligo/internal/order/repository"
	"github.com/maruki00/deligo/internal/order/router"
	"github.com/maruki00/deligo/internal/order/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	Repo       domain.OrderRepository
	Svc        domain.OrderService
	Handler    *handlers.OrderHandler
	HTTPServer *gin.Engine
	Consumer   *rabbitmq.Consumer
}

func Init(cfg *config.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {}, err
	}
	defer db.Close()

	mqClient, err := rabbitmq.NewRabbitMQClient(cfg.RabbitMQ.RabbitMQURL)
	if err != nil {
		return nil, func() {
			db.Close()
		}, err
	}
	defer mqClient.Close()

	consumer := rabbitmq.NewConsumer(mqClient)

	repo := repository.NewOrderRepository(db.DB)
	svc := service.NewOrderService(repo, mqClient)
	handler := handlers.NewOrderHandler(svc)

	return &App{
			Repo:       repo,
			Svc:        svc,
			Handler:    handler,
			HTTPServer: router.SetupRouter(handler),
			Consumer:   consumer,
		}, func() {
			db.Close()
			mqClient.Close()
		}, nil

}

func (a *App) Worder(ctx context.Context) error {
	return nil
}
