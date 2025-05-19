package middleware

import (
	"github.com/SiyovushAbdulloev/gophermart/internal/repository"
	jwt2 "github.com/SiyovushAbdulloev/gophermart/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(secret string, repository repository.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)
		userId := jwt2.CheckJWT(tokenString, secret)

		if userId == -1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong jwt token"})
			return
		}

		user, err := repository.GetUserById(userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
