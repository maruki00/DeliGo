package routes

import (
	"delivery/internal/auth/app"

	"github.com/gin-gonic/gin"
)

var AuthRouter = func(router *gin.Engine, app *app.App) {

	// repo := auth_infra_repository.NewAuthRepository(shared_configs.GetConfig())
	// controller := controllers.NewAuthController(app)

	// prefix := router.Group("/v1/auth")
	// _ = prefix.POST("/login", controller.Login)
	// _ = prefix.POST("/register", controller.Register)
	// _ = prefix.PATCH("/2f-confirm", controller.TwoFactoryConfirm)
	// _ = prefix.PATCH("/logout", controller.Logout)

}
