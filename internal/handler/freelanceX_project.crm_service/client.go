package handlers

import (
	"context"
	"net/http"
	"strings"
clientpb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_project.crm_service/client"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"github.com/Prototype-1/freelanceX_apigateway_service/pkg/jwt"
)

type ClientHandler struct {
	ClientClient clientpb.ClientServiceClient
}

func NewClientHandler(clientClient clientpb.ClientServiceClient) *ClientHandler {
	return &ClientHandler{ClientClient: clientClient}
}

func (h *ClientHandler) CreateClientHandler(c *gin.Context) {
	var req clientpb.CreateClientRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}
	claims, err := jwt.ParseAccessToken(tokenParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := claims.Role
	if role == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found in token"})
		return
	}

	sessionID := c.GetHeader("Session-Id")
if sessionID == "" {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Session-Id header is required"})
	return
}
	md := metadata.Pairs(
		"role", role,
		"Session-Id", sessionID,
		"user_id", claims.UserID,  
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := h.ClientClient.CreateClient(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res.Client)
}

func (h *ClientHandler) GetClientHandler(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}
	claims, err := jwt.ParseAccessToken(tokenParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := claims.Role
	if role != "client" && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only client or admin allowed"})
		return
	}

	sessionID := c.GetHeader("Session-Id")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session-Id header is required"})
		return
	}

	md := metadata.Pairs(
		"role", claims.Role,
		"Session-Id", sessionID,
		"user_id", claims.UserID,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := h.ClientClient.GetClient(ctx, &clientpb.GetClientRequest{ClientId: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Client)
}

func (h *ClientHandler) UpdateClientHandler(c *gin.Context) {
	id := c.Param("id")
	var req clientpb.UpdateClientRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ClientId = id

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}
	claims, err := jwt.ParseAccessToken(tokenParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := claims.Role
	if role != "client" && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only client or admin allowed"})
		return
	}

	sessionID := c.GetHeader("Session-Id")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session-Id header is required"})
		return
	}

	md := metadata.Pairs(
		"role", claims.Role,
		"Session-Id", sessionID,
		"user_id", claims.UserID,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := h.ClientClient.UpdateClient(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Client)
}

func (h *ClientHandler) DeleteClientHandler(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}
	claims, err := jwt.ParseAccessToken(tokenParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	role := claims.Role
	if role != "client" && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only client or admin allowed"})
		return
	}

	sessionID := c.GetHeader("Session-Id")
	if sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session-Id header is required"})
		return
	}

	md := metadata.Pairs(
		"role", claims.Role,
		"Session-Id", sessionID,
		"user_id", claims.UserID,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := h.ClientClient.DeleteClient(ctx, &clientpb.DeleteClientRequest{ClientId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": res.Status})
}

