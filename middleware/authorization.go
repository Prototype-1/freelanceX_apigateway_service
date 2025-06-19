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
        
        isWebSocket := c.GetHeader("Upgrade") == "websocket"
        log.Printf("Is WebSocket connection: %v", isWebSocket)
        
        var tokenStr string
        var sessionID string
        
        authHeader := c.GetHeader("Authorization")
        if authHeader != "" {
            parts := strings.Split(authHeader, "Bearer ")
            if len(parts) == 2 {
                tokenStr = parts[1]
            } else {
                log.Println("Invalid Authorization header format")
                c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
                return
            }
           sessionID = c.GetHeader("Session-Id")
if sessionID == "" {
    sessionID = c.GetHeader("session_id")
}
            log.Printf("Session ID from header: %s", sessionID)
        } else if isWebSocket {
    authHeader := c.GetHeader("Authorization")
    if authHeader != "" {
        parts := strings.Split(authHeader, "Bearer ")
        if len(parts) == 2 {
            tokenStr = parts[1]
        } else {
            log.Println("Invalid Authorization header format for WebSocket")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            return
        }
    }
    sessionID = c.GetHeader("Session-Id")
    displayToken := tokenStr
    if len(tokenStr) > 10 {
        displayToken = tokenStr[:10] + "...(truncated)"
    }
    log.Printf("WebSocket auth (from headers) - Token: %s, Session ID: %s", displayToken, sessionID)

    if tokenStr == "" || sessionID == "" {
        log.Println("Missing authentication headers for WebSocket")
        c.AbortWithStatusJSON(http.StatusUnauthorized, 
            gin.H{"error": "WebSocket connections require Authorization and session_id headers"})
        return
    }
} else {
          log.Println("Authorization header missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }
        claims, err := jwt.ParseAccessToken(tokenStr)
        if err != nil {
            log.Println("Token parsing error:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            return
        }
        
        log.Printf("User ID from token: %s", claims.UserID)
        
        if sessionID == "" {
            log.Println("Session ID missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session ID missing"})
            return
        }
        
        ctx := context.Background()
        key := fmt.Sprintf("session:%s", sessionID)
        storedUserID, err := redis.RedisClient.Get(ctx, key).Result()
        log.Printf("Stored User ID from Redis: %s", storedUserID)
        
        if err != nil {
            log.Println("Redis error:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session not found or expired. Please login again"})
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
