package handlers

import (
	"github.com/maruki00/deligo/internal/order/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *OrderHandler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.GlobalErrorHandler())

	orders := r.Group("/orders")
	{
		orders.POST("", handler.Create)
		orders.GET("/:id", handler.GetByID)
		orders.POST("/:id/confirm", handler.Confirm)
		orders.POST("/:id/accept", handler.Accept)
		orders.POST("/:id/complete", handler.Complete)
		orders.POST("/:id/cancel", handler.Cancel)
	}

	return r
}
