package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var FixedUsers = map[string]string{
	"user-1": "Alice",
	"user-2": "Bob",
	"user-3": "Charlie",
	"user-4": "Diana",
}

func UserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip validation for product routes and health check
		if c.Request.Method == "POST" && c.FullPath() == "/products" {
			c.Next()
			return
		}
		if c.Request.Method == "GET" && (c.FullPath() == "/products" || c.FullPath() == "/products/:id" || c.FullPath() == "/health") {
			c.Next()
			return
		}

		// For cart and invoice routes, validate user
		userID := c.GetHeader("X-User-Id")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "X-User-Id header is required",
			})
			c.Abort()
			return
		}

		if _, exists := FixedUsers[userID]; !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid user",
				"message": "User ID must be one of: user-1, user-2, user-3, user-4",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
