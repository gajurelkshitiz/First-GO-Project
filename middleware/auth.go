package middleware

import (
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			c.Abort()
			return
		}

		parts := strings.Split(h, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		// fmt.Println("Parts in Authorization: ", parts)

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Store user ID in context for next handlers
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
