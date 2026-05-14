package middleware

import (
	"CRUD_API_PROJ/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthHead := c.GetHeader("Authorization")
		if AuthHead == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		AuthHead = strings.TrimPrefix(AuthHead, "Bearer ")
		token, err := repository.ValidateJWT(AuthHead)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
