package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/payment/config"
	"github.com/maruki00/deligo/internal/payment/domain"
	"github.com/maruki00/deligo/internal/payment/handlers"
	"github.com/maruki00/deligo/internal/payment/rabbitmq"
	"github.com/maruki00/deligo/internal/payment/repository"
	"github.com/maruki00/deligo/internal/payment/router"
	"github.com/maruki00/deligo/internal/payment/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	Repo       domain.PaymentRepository
	Svc        domain.PaymentService
	Handler    *handlers.PaymentHandler
	HTTPServer *gin.Engine
	AMQPClient *rabbitmq.AMQPCliet
}

func Init(cfg *config.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {}, err
	}
	amqp, err := rabbitmq.NewAMQPClient(cfg)
	if err != nil {
		return nil, func() {
			db.Close()
		}, err
	}
	repo := repository.NewPaymentRepository(db.GetDB())
	svc := service.NewPaymentService(repo, amqp)
	handler := handlers.NewPaymentHandler(svc, "")

	return &App{
		Repo:       repo,
		Svc:        svc,
		Handler:    handler,
		HTTPServer: router.SetupRouter(handler),
		AMQPClient: amqp,
	}, nil, nil
}
