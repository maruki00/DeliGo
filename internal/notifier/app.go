package notifier

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/notifier/config"
	"github.com/maruki00/deligo/internal/notifier/domain"
	"github.com/maruki00/deligo/internal/notifier/handlers"
	"github.com/maruki00/deligo/internal/notifier/rabbitmq"
	"github.com/maruki00/deligo/internal/notifier/repository"
	"github.com/maruki00/deligo/internal/notifier/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	Repo        domain.NotifierRepository
	SVC         domain.NotifierService
	HTTPHandler *handlers.HTTPHandler
	WSHandler   *handlers.Hub
	AmqpClient  *rabbitmq.RabbitMQConnection
	HTTPServer  *gin.Engine
	///

}

func Init(cfg *config.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {}, err
	}

	repo := repository.NewNotifierRepository(db.GetDB())
	svc := service.NewNotifierService(repo)
	httpHandler := handlers.NewHTTPHandler(svc)
	hub := handlers.NewHub(svc)

	ampqClient := rabbitmq.NewRabbitMQConnection(cfg.RabbitMQ.URL, svc)

	return &App{
		Repo:        repo,
		SVC:         svc,
		HTTPHandler: httpHandler,
		WSHandler:   hub,
		AmqpClient:  ampqClient,
	}, nil, nil
}
