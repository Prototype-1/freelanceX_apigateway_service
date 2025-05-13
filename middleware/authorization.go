// package middleware

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"github.com/Prototype-1/freelanceX_apigateway_service/pkg/jwt"
// 	redis "github.com/Prototype-1/freelanceX_apigateway_service/pkg/redis" 
// 	"github.com/gin-gonic/gin"
// 	"log"
// )

// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         log.Println("--- Request Headers ---")
//         for k, v := range c.Request.Header {
//             log.Printf("%s: %v", k, v)
//         }

//          isWebSocket := c.GetHeader("Upgrade") == "websocket"
//         log.Printf("Is WebSocket connection: %v", isWebSocket)
        
//         authHeader := c.GetHeader("Authorization")
//         if authHeader == "" {
//             log.Println("Authorization header missing")
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
//             return
//         }
        
//         parts := strings.Split(authHeader, "Bearer ")
//         if len(parts) != 2 {
//             log.Println("Invalid Authorization header format")
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
//             return
//         }
        
//         tokenStr := parts[1]
//         claims, err := jwt.ParseAccessToken(tokenStr)
//         if err != nil {
//             log.Println("Token parsing error:", err)
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
//             return
//         }
        
//         sessionID := c.GetHeader("session_id")
//         log.Printf("Session ID from header: %s", sessionID)
//         log.Printf("User ID from token: %s", claims.UserID)
        
//         if sessionID == "" {
//             log.Println("Session ID header missing")
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session ID header missing"})
//             return
//         }
        
//         ctx := context.Background()
//         key := fmt.Sprintf("session:%s", sessionID)
//         storedUserID, err := redis.RedisClient.Get(ctx, key).Result()
//         log.Printf("Stored User ID from Redis: %s", storedUserID)
        
//         if err != nil {
//             log.Println("Redis error:", err)
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session not found or expired. Please login again"})
//             return
//         }
        
//         if storedUserID != claims.UserID {
//             log.Println("User ID mismatch")
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session: User ID mismatch"})
//             return
//         }
        
//         c.Set("user_id", claims.UserID)
//         c.Set("role", claims.Role)
//         c.Set("session_id", sessionID)
//         log.Printf("Authenticated user %s with role %s", claims.UserID, claims.Role)
        
//         c.Next()
//     }
// }

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
        
        // Check if this is a WebSocket connection
        isWebSocket := c.GetHeader("Upgrade") == "websocket"
        log.Printf("Is WebSocket connection: %v", isWebSocket)
        
        // Get token from Authorization header or query parameter for WebSockets
        var tokenStr string
        var sessionID string
        
        // Standard header-based auth
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
            
            // Get session ID from header
            sessionID = c.GetHeader("session_id")
            log.Printf("Session ID from header: %s", sessionID)
        } else if isWebSocket {
            tokenStr = c.Query("token")
            sessionID = c.Query("session_id")

          displayToken := tokenStr
    if len(tokenStr) > 10 {
        displayToken = tokenStr[:10] + "...(truncated)"
    }
    log.Printf("WebSocket auth - Token: %s, Session ID: %s", displayToken, sessionID)

            if tokenStr == "" || sessionID == "" {
                log.Println("Missing authentication parameters for WebSocket")
                c.AbortWithStatusJSON(http.StatusUnauthorized, 
                    gin.H{"error": "WebSocket connections require token and session_id query parameters"})
                return
            }
        } else {
            log.Println("Authorization header missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }
        
        // Parse and validate the JWT token
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
        
        // Verify session in Redis
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
        
        // Authentication successful - set context values
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Set("session_id", sessionID)
        log.Printf("Authenticated user %s with role %s", claims.UserID, claims.Role)
        
        c.Next()
    }
}