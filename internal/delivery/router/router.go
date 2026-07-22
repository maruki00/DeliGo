package router

import (
	"github.com/maruki00/deligo/internal/delivery/handlers"
	"github.com/maruki00/deligo/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.DeliveryHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger())

	api := r.Group("/api/v1")
	{
		api.POST("/couriers", h.RegisterCourier)
		api.POST("/tracking/ping", h.PingLocation)

		orders := api.Group("/orders/:order_id")
		{
			orders.POST("/accept", h.AcceptOrder)
			orders.POST("/arrive", h.ArriveAtRestaurant)
			orders.POST("/pickup", h.StartDelivery)
			orders.POST("/complete", h.CompleteDelivery)
		}
	}
	return r
}
