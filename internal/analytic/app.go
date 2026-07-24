package analytic

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/analytic/config"
	"github.com/maruki00/deligo/internal/analytic/domain"
	"github.com/maruki00/deligo/internal/analytic/handlers"
	"github.com/maruki00/deligo/internal/analytic/rabbitmq"
	"github.com/maruki00/deligo/internal/analytic/repository"
	"github.com/maruki00/deligo/internal/analytic/router"
	"github.com/maruki00/deligo/internal/analytic/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	ARepo      domain.AnalyticRepository
	FRepo      domain.FeedbackRepository
	FSvc       domain.FeedbackService
	Publisher  domain.EventPublisher
	Handler    *handlers.FeedbackHandler
	HTTPServer *gin.Engine
}

func Init(cfg *config.Config) (*App, func(), error) {
	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {}, err
	}

	aRepo := repository.NewAnalyticRepository(db.DB)
	fRepo := repository.NewFeedbackRepository(db.DB)

	_, amqpChan := rabbitmq.InitRabbitMQ(cfg)
	_ = rabbitmq.NewEventConsumer(amqpChan, cfg, aRepo)
	publisher := rabbitmq.NewEventPublisher(amqpChan, cfg)

	fSvc := service.NewFeedbackService(fRepo, aRepo, nil)
	handler := handlers.NewFeedbackHandler(fSvc)
	return &App{
		ARepo:      aRepo,
		FRepo:      fRepo,
		FSvc:       fSvc,
		Publisher:  publisher,
		Handler:    handler,
		HTTPServer: router.SetupRouter(handler),
	}, nil, nil
}
