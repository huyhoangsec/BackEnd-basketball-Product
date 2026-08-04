package middleware

import (
	"net/http"
	"strings"

	"backend-ocean-basketball/config"
	"backend-ocean-basketball/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString, cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check roles if specified
		if len(requiredRoles) > 0 {
			hasRole := false
			for _, role := range requiredRoles {
				if claims.Role == role {
					hasRole = true
					break
				}
			}
			if !hasRole {
				c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				c.Abort()
				return
			}
		}

		// Set context variables
		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)
		c.Next()
	}
}
