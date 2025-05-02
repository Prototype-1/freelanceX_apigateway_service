package handler

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
authPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/auth"
)

func Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"` 
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.Register(ctx, &authPb.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		if err.Error() == "this email is already registered" {
			c.JSON(http.StatusConflict, gin.H{"error": "This email is already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": res.AccessToken})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.Login(ctx, &authPb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": res.AccessToken,
		"session_id":   res.SessionId,
		"user_id":      res.UserId,
	})
}

func OAuth(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Provider string `json:"provider" binding:"required"`
		OAuthId  string `json:"oauth_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.OAuthLogin(ctx, &authPb.OAuthRequest{
		Email:         req.Email,
		OauthProvider: req.Provider,
		OauthId:       req.OAuthId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !res.IsRoleSelected {
		c.JSON(http.StatusOK, gin.H{
			"user_id":          res.UserId,
			"is_role_selected": false,
			"message":          "Role selection required",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          res.Message,
		"access_token":     res.AccessToken,
		"session_id":       res.SessionId,
		"user_id":          res.UserId,
		"is_role_selected": res.IsRoleSelected,
		"name":             res.Name,
		"email":            res.Email,
		"role":             res.Role,
	})
}

func SelectRole(c *gin.Context) {
	var req struct {
		UserId string `json:"user_id" binding:"required"`
		Role   string `json:"role" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.SelectRole(ctx, &authPb.SelectRoleRequest{
		UserId: req.UserId,
		Role:   req.Role,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": res.Message,
	})
}

func GetMe(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-Session-ID header"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.GetMe(ctx, &authPb.SessionRequest{
		SessionId: sessionID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":              res.Id,
			"name":            res.Name,
			"email":           res.Email,
			"role":            res.Role,
			"is_role_selected": res.IsRoleSelected,
		},
	})
}

func Logout(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing session ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := client.AuthClient.Logout(ctx, &authPb.LogoutRequest{
		SessionId: sessionID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

