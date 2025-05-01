package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/Prototype-1/freelanceX_user_service/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type contextKey string
const userIDKey contextKey = "userID"


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenStr := parts[1]
		claims, err := jwt.ParseAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		sessionID := c.GetHeader("X-Session-ID")
		if sessionID == "" || !jwt.ValidateSession(sessionID, claims.UserID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), userIDKey, claims.UserID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
