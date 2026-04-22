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
		taskID := c.Param("id")
		AuthHead = strings.TrimPrefix(AuthHead, "Bearer ")
		isValid, err := repository.ValidateToken(AuthHead, taskID)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
