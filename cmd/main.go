package main

import (
	"github.com/maruki00/deligo/internal/iam/infra/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=172.18.0.3 port=5432 user=admin password=admin dbname=github.com/maruki00/deligo sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("migration failed: " + err.Error())
	}

	fmt.Println("Migration successful!")
}
