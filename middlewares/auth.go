package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := os.Getenv("TOKEN")
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda Belum login"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization tidak valid"})
			c.Abort()
			return
		}

		receivedToken := strings.TrimPrefix(authorization, "Bearer ")

		if receivedToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
