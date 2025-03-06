package main

import (
	"delivery/cmd/courier/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configs.GetConfig("")
	router := gin.New()
	router.Run(fmt.Sprintf("%s:%s", config.RestServer.Host, config.RestServer.Port))
}
