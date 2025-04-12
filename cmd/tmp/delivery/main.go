package main

import (
	auth_routes "delivery/internal/auth/Application/Routes"
	order_routes "delivery/internal/order/Application/Routes"
	prouduct_routes "delivery/internal/product/Application/Routes"
	shared_configs "delivery/internal/shared/Application/Configs"
	shareddb "delivery/internal/shared/infra/DB"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := shared_configs.GetConfig()
	db := shareddb.NewMysqlDB_GORM(config)
	router := gin.Default()
	go auth_routes.AuthRouter(router, db)
	prouduct_routes.ProductRouter(router, db)
	order_routes.OrderRouter(router, db)

	router.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
