package router

import (
	"github.com/maruki00/deligo/internal/analytic/handlers"
	"github.com/maruki00/deligo/internal/analytic/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.FeedbackHandler) *gin.Engine {
	r := gin.New()

	// Inject architectural custom tracking components
	r.Use(middleware.StructuredLogger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/feedbacks", h.CreateFeedback)
		v1.GET("/analytics/:product_id", h.GetProductAnalytics)
	}

	return r
}
