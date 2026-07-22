package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"user-profile-service/internal/config"
	"user-profile-service/internal/handlers"
	"user-profile-service/internal/middleware"
	"user-profile-service/internal/rabbitmq"
	"user-profile-service/internal/repository"
	"user-profile-service/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database link setup failed completely: %v", err)
	}

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Casbin storage driver attachment failed: %v", err)
	}

	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "r.sub == p.sub && keyMatch2(r.obj, p.obj) && r.act == p.act")

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("Casbin authorization engine initialization failed: %v", err)
	}
	_ = enforcer.LoadPolicy()
	_ = middleware.SeedDefaultRules(enforcer)

	publisher, err := rabbitmq.NewPublisher(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("RabbitMQ message production system link failed: %v", err)
	}
	defer publisher.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, publisher, cfg)
	userHandler := handlers.NewUserHandler(userService, enforcer)

	router := gin.Default()

	// Unprotected operational boundaries
	router.POST("/api/v1/auth/register", userHandler.Register)
	router.POST("/api/v1/auth/login", userHandler.Login)

	// Protected operational boundaries
	secureAPI := router.Group("/api/v1")
	secureAPI.Use(middleware.SecurityEngine(enforcer, cfg))
	{
		secureAPI.GET("/users/:id", userHandler.GetProfile)
		secureAPI.PUT("/users/:id", userHandler.UpdateProfile)
		secureAPI.DELETE("/users/:id", userHandler.Delete)
		secureAPI.POST("/users/:id/ban", userHandler.Ban)
		secureAPI.POST("/permissions", userHandler.AssignPermission)
	}

	server := &http.Server{Addr: ":" + cfg.Port, Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Network socket layer terminated unexpectedly: %v", err)
		}
	}()
	log.Printf("Unified IAM and User Profile microservice launched securely on port: %s", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
	log.Println("Operational engine shut down safely without resource leaks.")
}
