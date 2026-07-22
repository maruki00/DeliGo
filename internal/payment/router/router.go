package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/payment/handlers"
	"github.com/maruki00/deligo/internal/payment/middleware"
)

func SetupRouter(h *handlers.PaymentHandler) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	paymentRoutes := r.Group("/payments")
	{
		paymentRoutes.POST("/charge", h.Charge)
		paymentRoutes.POST("/refund", h.Refund)
		paymentRoutes.POST("/webhook", h.Webhook)
	}

	return r
}
