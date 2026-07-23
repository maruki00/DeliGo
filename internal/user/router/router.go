package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/user/handlers"
)

var SetupRouter = func(handler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/api/v1/auth/register", handler.Register)
	router.POST("/api/v1/auth/login", handler.Login)
	router.POST("/api/v1/auth", handler.Auth)

	secureAPI := router.Group("/api/v1")
	// secureAPI.Use(middleware.SecurityEngine(enforcer, cfg))
	// {
	secureAPI.GET("/users/:id", handler.GetUser)
	secureAPI.PUT("/users/:id", handler.UpdateUser)
	secureAPI.DELETE("/users/:id", handler.Delete)
	secureAPI.POST("/users/:id/ban", handler.Ban)
	// secureAPI.POST("/permissions", 		app.Handler.AssignPermission)
	// }

	return router

}
