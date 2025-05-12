package utils

import (
    "context"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc/metadata"
    "log"
)

func InjectMetadataFromGin(c *gin.Context) context.Context {
    role, exists := c.Get("role")
    if !exists {
        log.Println("Role not found in Gin context")
        return metadata.NewOutgoingContext(c.Request.Context(), metadata.Pairs())
    }
    
    roleStr, ok := role.(string)
    if !ok {
        log.Printf("Role is not a string: %T", role)
        return metadata.NewOutgoingContext(c.Request.Context(), metadata.Pairs())
    }
    
    log.Printf("Injecting role metadata: %s", roleStr)
    
    md := metadata.Pairs("role", roleStr)
    if userID, exists := c.Get("user_id"); exists {
        if userIDStr, ok := userID.(string); ok {
            md = metadata.Join(md, metadata.Pairs("user_id", userIDStr))
        }
    }
    
    if sessionID, exists := c.Get("session_id"); exists {
        if sessionIDStr, ok := sessionID.(string); ok {
            md = metadata.Join(md, metadata.Pairs("session_id", sessionIDStr))
        }
    }
    
    return metadata.NewOutgoingContext(c.Request.Context(), md)
}