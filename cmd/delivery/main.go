package main

import (
	"delivery/internal/config"
	"delivery/internal/handlers"
	"delivery/internal/models"
	"delivery/internal/rabbitmq"
	"delivery/internal/repository"
	"delivery/internal/router"
	"delivery/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the DB database: %v", err)
	}

	err = db.AutoMigrate(&models.Courier{}, &models.OrderCourier{}, &models.OrderTracking{})
	if err != nil {
		log.Fatalf("Failed database auto-migration: %v", err)
	}

	repo := repository.NewDeliveryRepository(db)
	svc := service.NewDeliveryService(repo)
	handler := handlers.NewDeliveryHandler(svc)

	consumer := rabbitmq.NewRabbitMQConsumer(cfg.RabbitMQURL, svc)
	consumer.Start()
	defer consumer.Close()

	r := router.SetupRouter(handler)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server initializing execution on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen exception encountered: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down HTTP execution context cleanly...")
}
