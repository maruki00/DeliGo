package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const UserContextKey = "owner_id"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := strings.TrimSpace(c.GetHeader("X-User-ID"))
		if userID == "" {
			// Requirements explicitly noted returning 411 Length Required / Unauthorized variant context
			c.AbortWithStatusJSON(http.StatusLengthRequired, gin.H{"error": "X-User-ID header is missing or empty"})
			return
		}
		c.Set(UserContextKey, userID)
		c.Next()
	}
}
