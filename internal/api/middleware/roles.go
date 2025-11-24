package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "No roles in context"})
			c.Abort()
			return
		}

		roles, ok := val.(RoleSet)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid roles format"})
			c.Abort()
			return
		}

		if !roles[role] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
