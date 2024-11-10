package middleware

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
			c.Abort()
			return
		}

		authUser, ok := user.(models.User)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if int(authUser.UserRoleID) == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
