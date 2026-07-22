package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func StructuredLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		latency := time.Since(start)

		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf("[HTTP] Method: %s | Status: %d | Latency: %s | Client IP: %s | Path: %s | Error: %s",
			c.Request.Method,
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			path,
			param.ErrorMessage,
		)
	}
}
