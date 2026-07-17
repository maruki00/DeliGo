package main

import (
	"log"

	"github.com/maruki00/deligo/internal/profile/app"
	profilecfg "github.com/maruki00/deligo/internal/profile/configs"
	profileapi "github.com/maruki00/deligo/internal/profile/trensport/api"
)

func main() {
	cfg := &profilecfg.Config{}
	cfg.Postgres.DSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	instance, cleanup, err := app.InitApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize profile service: %v", err)
	}
	defer cleanup()

	router := profileapi.NewRouter(instance)
	log.Printf("profile service initialized")
	if err := router.Run(":8091"); err != nil {
		log.Fatalf("failed to serve profile service: %v", err)
	}
}
