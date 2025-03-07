package routes

import "github.com/gin-gonic/gin"

func ShopRouter(router *gin.Engine) {
	shop := router.Group("/restaurant")

	_ = shop.POST("/open", nil)
	_ = shop.POST("/close", nil)

	employee := shop.Group("/employee")
	employee.POST("/add", nil)
	employee.POST("/delete", nil)
	employee.POST("/get", nil)
	employee.POST("/search", nil)

	menu := shop.Group("/menu")
	menu.POST("/add", nil)
	menu.POST("/delete", nil)
	menu.POST("/edit", nil)
	menu.POST("/get", nil)
	menu.POST("/search", nil)

}
