package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/user/config"
	"github.com/maruki00/deligo/internal/user/domain"
	"github.com/maruki00/deligo/internal/user/handlers"
	"github.com/maruki00/deligo/internal/user/rabbitmq"
	"github.com/maruki00/deligo/internal/user/repository"
	"github.com/maruki00/deligo/internal/user/router"
	"github.com/maruki00/deligo/internal/user/service"
	pkgPostgres "github.com/maruki00/deligo/pkg/postgres"
)

type App struct {
	DB         *pkgPostgres.PGHandler
	Repo       domain.UserRepository
	Svc        domain.UserService
	Handler    *handlers.UserHandler
	HTTPServer *gin.Engine
}

func Init(cfg *config.Config) (*App, func(), error) {

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {
			// cancel()
		}, err
	}
	publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQ.RabbitMQURL)
	if err != nil {
		return nil, func() {
			db.Close()
			// cancel()
		}, err
	}
	repo := repository.NewUserRepository(db.DB)
	svc := service.NewUserService(repo, publisher, cfg)
	hdler := handlers.NewUserHandler(svc)

	// adapter, err := gormadapter.NewAdapterByDB(db.GetDB())
	// if err != nil {
	// 	return nil, func() {
	// 		db.Close()
	// 		publisher.Close()
	// 	}, fmt.Errorf("Casbin storage driver attachment failed: %v", err)
	// }

	// m := model.NewModel()
	// m.AddDef("r", "r", "sub, obj, act")
	// m.AddDef("p", "p", "sub, obj, act")
	// m.AddDef("e", "e", "some(where (p.eft == allow))")
	// m.AddDef("m", "m", "r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act")

	// enforcer, err := casbin.NewEnforcer(m, adapter)
	// if err != nil {
	// 	return nil, func() {
	// 		db.Close()
	// 		publisher.Close()
	// 	}, fmt.Errorf("Casbin authorization engine initialization failed: %v", err)
	// }
	// _ = enforcer.LoadPolicy()
	// _ = middleware.SeedDefaultRules(enforcer)

	return &App{
			DB:         db,
			Repo:       repo,
			Svc:        svc,
			Handler:    hdler,
			HTTPServer: router.SetupRouter(hdler),
		}, func() {
			db.Close()
			publisher.Close()
			// cancel()
		}, nil
}
