package handler

import (
	"context"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
authPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/auth"
"github.com/Prototype-1/freelanceX_apigateway_service/pkg/oauth"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, password, and role
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User registration payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
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

	c.JSON(http.StatusCreated, gin.H{
		"message": res.Message,
		"user_id": res.UserId,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
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

// OAuth godoc
// @Summary Login with OAuth (Google etc.)
// @Description Authenticate user via OAuth provider
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.OAuthRequest true "OAuth request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/oauth [post]
func OAuth(c *gin.Context) {
	var req struct {
		Code     string `json:"code" binding:"required"`
		Provider string `json:"provider" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userProfile, err := oauth.ExchangeAuthorizationCodeForToken(req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oAuthReq := &authPb.OAuthRequest{
		OauthProvider: req.Provider,
		Code:          req.Code, 
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.AuthClient.OAuthLogin(ctx, oAuthReq)
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
		"name":             userProfile.Name,
		"email":            userProfile.Email,
		"role":             res.Role,
	})
}

// SelectRole godoc
// @Summary Select user role
// @Description Allow user to select a role after OAuth login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.SelectRoleRequest true "Select role request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/select-role [post]
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

// GetMe godoc
// @Summary Get current user info
// @Description Fetch user information using access token and session context
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/me [get]
func GetMe(c *gin.Context) {
    userID := c.GetString("user_id")
    sessionID := c.GetString("session_id") 

    if userID == "" || sessionID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: missing user context"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res, err := client.AuthClient.GetMe(ctx, &authPb.SessionRequest{
        Token:     c.GetHeader("Authorization")[7:], 
        SessionId: sessionID,
        UserId:    userID, 
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": gin.H{
            "id":               res.Id,
            "name":             res.Name,
            "email":            res.Email,
            "role":             res.Role,
            "is_role_selected": res.IsRoleSelected,
        },
    })
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate the current session
// @Tags Auth
// @Produce json
// @Param session_id header string true "Session ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/logout [post]
func Logout(c *gin.Context) {
	sessionID := c.GetHeader("session_id")
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

