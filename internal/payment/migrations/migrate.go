package migrations

import (
	"log"

	"github.com/maruki00/deligo/internal/payment/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migration completed successfully.")
}
