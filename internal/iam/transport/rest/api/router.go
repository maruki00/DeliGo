package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maruki00/deligo/internal/iam/transport/rest/handler"
)

var router = func(router *gin.Engine) {
	var api = router.Group("api/v1/")
	var authz = api.Group("authz")

	authz.POST("/create-policy", handler.CreatePolicy)
	authz.POST("/create-policy", handler.CreatePolicy)
	authz.POST("/create-policy", handler.CreatePolicy)
	authz.POST("/create-policy", handler.CreatePolicy)
}
