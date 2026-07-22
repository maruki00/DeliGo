package router

import (
	"github.com/maruki00/deligo/internal/notifier/handlers"
	"github.com/maruki00/deligo/internal/notifier/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(httpH *handlers.HTTPHandler, wsH *handlers.Hub) *gin.Engine {
	r := gin.Default()

	// WebSocket real-time system pipeline mapping (unprotected for web standards fallback/custom handshakes)
	r.GET("/ws", wsH.HandleWebSocket)

	// Protected Application Interface Routes Setup
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/notifications", httpH.GetNotifications)
		api.PATCH("/notifications/:id/read", httpH.MarkAsRead)
		api.GET("/chat/history/:order_id", httpH.GetChatHistory)
		api.POST("/files/upload", httpH.UploadFile)
	}

	return r
}
