package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"github.com/Prototype-1/freelanceX_apigateway_service/pkg/jwt"
	"github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis" 
	"github.com/gin-gonic/gin"
)

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
		ctx := context.Background()
		key := fmt.Sprintf("session:%s", sessionID)
		exists, err := pkg.RedisClient.Exists(ctx, key).Result()
		if err != nil || exists == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session not found or expired"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("session_id", sessionID)
		c.Next()
	}
}
