package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"delivery/internal/product/userGetway/controllers"
)

var ProductRouter = func(router *gin.Engine, db *gorm.DB) {

	controller := controllers.NewProductController(db)

	prefix := router.Group("/api/product")
	_ = prefix.POST  ("/insert", controller.Insert)
	_ = prefix.POST  ("/update", controller.Update)
	_ = prefix.DELETE("/delete", controller.Delete)
	_ = prefix.POST  ("/search", controller.Search)
	_ = prefix.POST  ("/get", controller.GetProduct)
	_ = prefix.POST  ("/multiple", controller.MultipleProducts)
}


post:"/v0/product/insert",
body: "*"
post:"/v0/product/update",
body: "*"
delete:"/v0/product/delete",
body: "*"
post:"/v0/product/search",
body: "*"
post:"/v0/product/get",
body: "*"
post:"/v0/product/multiple",
body: "*"