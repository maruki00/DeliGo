package main

import (
	"delivery/cmd/shop/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configs.GetConfig(".")
	router := gin.Default()
	router.Run(fmt.Sprintf("%s:%s", config.RestServer.Host, config.RestServer.Port))
}
