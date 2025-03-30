package utils

import (
	"github.com/gin-gonic/gin"
)

func Success(context *gin.Context, httpStatus int, message string, data map[string]any) {
	context.JSON(httpStatus, gin.H{
		"status":  httpStatus,
		"message": message,
		"data":    data,
	})
}

func Error(context *gin.Context, httpStatus int, message string, err string) {
	context.JSON(httpStatus, gin.H{
		"status":  httpStatus,
		"message": message,
		"error":   err,
	})
}
