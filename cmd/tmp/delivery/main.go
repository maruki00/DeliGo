package main

import (
	auth_routes "deligo/internal/auth/Application/Routes"
	order_routes "deligo/internal/order/Application/Routes"
	prouduct_routes "deligo/internal/product/Application/Routes"
	shared_configs "deligo/internal/shared/Application/Configs"
	shareddb "deligo/internal/shared/infra/DB"
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
