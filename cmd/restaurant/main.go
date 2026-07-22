package main

import (
	"fmt"
	"os"

	"github.com/maruki00/deligo/cmd/restaurant/config"
	"github.com/maruki00/deligo/internal/restaurant"
	"github.com/maruki00/deligo/internal/restaurant/repositories"
)

func main() {

	rootPath, _ := os.Getwd()
	println(rootPath + "/../../migrations/000001_init_restaurant_schema.up.sql")
	cfgPath := rootPath + "/config/config.yaml"
	cfg, err := config.NewCfg(cfgPath)
	if err != nil {
		panic(err)
	}

	app, clean, err := restaurant.Iinit(cfg)
	if err != nil {
		clean()
		panic(err)
	}
	defer clean()
	repositories.RunDatabaseMigrations(cfg.DB.Dsn, "file:///home/user/dev/DeliGo/migrations")

	app.HTTPServer.Run(fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port))
}
