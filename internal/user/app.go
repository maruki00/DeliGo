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
	UserRepo   domain.UserRepository
	AuthRepo   domain.AuthRepository
	Svc        domain.UserService
	Handler    *handlers.UserHandler
	HTTPServer *gin.Engine
}

func Init(cfg *config.Config) (*App, func(), error) {

	db, err := pkgPostgres.NewDB(cfg.DB.DSN)
	if err != nil {
		return nil, func() {
		}, err
	}
	publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQ.RabbitMQURL)
	if err != nil {
		return nil, func() {
			db.Close()
		}, err
	}
	uRepo := repository.NewUserRepository(db.DB)
	authRepo := repository.NewAuthRepository(db.DB)
	svc := service.NewUserService(uRepo, authRepo, publisher, cfg)
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
			UserRepo:   uRepo,
			AuthRepo:   authRepo,
			Svc:        svc,
			Handler:    hdler,
			HTTPServer: router.SetupRouter(hdler),
		}, func() {
			db.Close()
			publisher.Close()
			// cancel()
		}, nil
}
