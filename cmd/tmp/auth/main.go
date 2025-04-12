package main

import (
	"delivery/cmd/auth/configs"
	"delivery/internal/auth/app"
	"delivery/internal/auth/app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {

	_, err := maxprocs.Set()
	if err != nil {
		panic(err)
	}

	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	app, err := app.InitApp(cfg)

	server := gin.Default()
	routes.AuthRouter(server, app)
	server.Run(fmt.Sprintf("%s:%s", cfg.RestServer.Host, cfg.RestServer.Port))
}
