package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "already authenticated"})
			return
		}

		c.Next()
	}
}
