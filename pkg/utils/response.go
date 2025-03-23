package utils

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, message string, result any) {
	ctx.JSON(
		200,
		gin.H{
			"message": message,
			"status":  200,
			"result":  result,
		},
	)
}

func Error(ctx *gin.Context, message string, result any) {
	ctx.JSON(
		200,
		gin.H{
			"message": message,
			"status":  400,
			"result":  result,
		},
	)
}
