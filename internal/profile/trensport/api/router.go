package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/profile/app"
	"github.com/maruki00/deligo/internal/profile/trensport/handler"
)

func NewRouter(instance *app.App) *gin.Engine {
	r := gin.Default()
	h := handler.New(instance)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	profiles := r.Group("/profiles")
	profiles.POST("", h.CreateProfile)
	profiles.GET(":id", h.GetProfile)
	profiles.PUT(":id", h.UpdateProfile)
	profiles.PATCH(":id/avatar", h.UpdateAvatar)
	profiles.DELETE(":id", h.DisableProfile)

	return r
}
