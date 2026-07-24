package database

import (
	"log"
	"time"

	"github.com/maruki00/deligo/internal/analytic/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Models struct {
	Feedback interface{}
	Analytic interface{}
}

func InitDB(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	// Simple retry mechanism for startup resilience in containerized stacks
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying in 3 seconds... (%d/5)", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Fatal: Could not connect to database after retries: %v", err)
	}

	log.Println("Database connection established successfully.")
	return db
}
