package utils

import (
	"context"
"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func InjectMetadataFromGin(c *gin.Context) context.Context {
	roleVal, _ := c.Get("role")
	role, _ := roleVal.(string)

	md := metadata.Pairs("role", role)
	return metadata.NewOutgoingContext(c.Request.Context(), md)
}
