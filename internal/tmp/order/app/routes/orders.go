package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var OrderRouter = func(router *gin.Engine, db *gorm.DB) {
	//controller := order_usergateway_controllers.NewOrderController()
	order := router.Group("/order/")
	_ = order.POST(" /create", nil)
	_ = order.POST(" /cancel", nil)
	_ = order.POST(" /history/{id}", nil)
	_ = order.POST(" /pickup-confirm/{id}", nil)
	_ = router.POST("/chat", nil)
}
