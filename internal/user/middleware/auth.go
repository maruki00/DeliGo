package middleware

import (
	"net/http"
	"strings"

	"github.com/maruki00/deligo/internal/user/config"
	"github.com/maruki00/deligo/internal/user/domain"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SecurityEngine(enforcer *casbin.Enforcer, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing signature authorization carrier"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid carrier authentication type format"})
			return
		}

		claims := &domain.JWTClaims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Crypto parsing check failed for provided payload signature"})
			return
		}

		// Hydrate session context state variables
		c.Set("user_id", claims.Subject)
		c.Set("user_role", claims.Role)

		// Evaluate request against live policy enforcer
		obj := c.Request.URL.Path
		act := c.Request.Method

		allowed, err := enforcer.Enforce(claims.Role, obj, act)
		if err != nil || !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Security clearance verification failed for current context"})
			return
		}

		c.Next()
	}
}

func SeedDefaultRules(enforcer *casbin.Enforcer) error {
	if hasRules := enforcer.HasPolicy("admin", "/api/v1/users/:id/ban", "POST"); !hasRules {
		rules := [][]string{
			// Admin global privileges
			{"admin", "/api/v1/users/:id", "GET"},
			{"admin", "/api/v1/users/:id", "PUT"},
			{"admin", "/api/v1/users/:id", "DELETE"},
			{"admin", "/api/v1/users/:id/ban", "POST"},
			{"admin", "/api/v1/permissions", "POST"},

			// Regular resource boundaries
			{"customer", "/api/v1/users/:id", "GET"},
			{"customer", "/api/v1/users/:id", "PUT"},
			{"courier", "/api/v1/users/:id", "GET"},
			{"courier", "/api/v1/users/:id", "PUT"},
		}
		_, err := enforcer.AddPolicies(rules)
		if err != nil {
			return err
		}
		return enforcer.SavePolicy()
	}
	return nil
}
