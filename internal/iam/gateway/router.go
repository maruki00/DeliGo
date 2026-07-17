package api

import (
	"github.com/gin-gonic/gin"
)

var userRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var user = api.Group("user")

	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
	user.POST("/create-policy", nil)
}

var identityRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var identity = api.Group("identity")

	identity.POST("add-policy", nil)
	identity.POST("add-group-policy", nil)
	identity.POST("remove-policy", nil)
	identity.POST("remove-group-policy", nil)
	identity.POST("affect-permission", nil)

}

var authRouter = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var auth = api.Group("auth")

	auth.POST("login", nil)
}

/*

user/aprove
user/ban
user/get
user/create



identity/add-policy
identity/add-group-policy
identity/remove-policy
identity/remove-group-policy
identity/affect-permission

auth/login
auth/grante-access //grpc

*/
