package main

import (
	"delivery/cmd/user/configs"
	"delivery/internal/user/app"

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

}
