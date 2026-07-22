package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/restaurant/handlers"
	"github.com/maruki00/deligo/internal/restaurant/middlewares"
)

func SetupRouter(h *handlers.CatalogHandler) *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		// Public Routes
		apiV1.GET("/restaurants/:id/menu", h.GetRestaurantMenu)

		// Owner Protected Routes
		protected := apiV1.Group("")
		protected.Use(middlewares.AuthMiddleware())
		{
			protected.POST("/restaurants", h.CreateRestaurant)
			protected.PATCH("/restaurants/:id/status", h.PatchRestaurantStatus)
			protected.POST("/restaurants/:id/products", h.AddProduct)
			protected.PUT("/products/:product_id", h.UpdateProduct)
			protected.DELETE("/products/:product_id", h.DeleteProduct)
		}
	}

	return r
}
