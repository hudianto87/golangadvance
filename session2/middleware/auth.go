package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, pass, ok := c.Request.BasicAuth()

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			c.Abort()
			return
		}

		const (
			expectedUsername = "admin"
			expectedPassword = "admin123"
		)

		isvalid := (username == expectedUsername) && (pass == expectedPassword)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			c.Abort()
			return
		}

		c.Next()
	}
}
