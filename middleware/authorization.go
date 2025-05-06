package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"github.com/Prototype-1/freelanceX_apigateway_service/pkg/jwt"
	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis" 
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("--- Request Headers ---")
        for k, v := range c.Request.Header {
            log.Printf("%s: %v", k, v)
        }
        
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            log.Println("Authorization header missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }
        
        parts := strings.Split(authHeader, "Bearer ")
        if len(parts) != 2 {
            log.Println("Invalid Authorization header format")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            return
        }
        
        tokenStr := parts[1]
        claims, err := jwt.ParseAccessToken(tokenStr)
        if err != nil {
            log.Println("Token parsing error:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }
        
        sessionID := c.GetHeader("session_id")
        log.Printf("Session ID from header: %s", sessionID)
        log.Printf("User ID from token: %s", claims.UserID)
        
        if sessionID == "" {
            log.Println("Session ID header missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session ID header missing"})
            return
        }
        
        ctx := context.Background()
        key := fmt.Sprintf("session:%s", sessionID)
        storedUserID, err := redis.RedisClient.Get(ctx, key).Result()
        log.Printf("Stored User ID from Redis: %s", storedUserID)
        
        if err != nil {
            log.Println("Redis error:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session not found or expired"})
            return
        }
        
        if storedUserID != claims.UserID {
            log.Println("User ID mismatch")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session: User ID mismatch"})
            return
        }
        
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Set("session_id", sessionID)
        log.Printf("Authenticated user %s with role %s", claims.UserID, claims.Role)
        
        c.Next()
    }
}

