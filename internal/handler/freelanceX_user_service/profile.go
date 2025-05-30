package handler

import (
	"net/http"
"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	profilePb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/profile"
	"github.com/gin-gonic/gin"
		"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service/utils"
)

func CreateProfile(c *gin.Context) {
	var req profilePb.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := utils.InjectMetadataFromGin(c)

	res, err := client.ProfileClient.CreateProfile(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateProfile(c *gin.Context) {
	var req profilePb.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := utils.InjectMetadataFromGin(c)

	res, err := client.ProfileClient.UpdateProfile(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetProfile(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
	ctx := utils.InjectMetadataFromGin(c)

	req := &profilePb.GetProfileRequest{UserId: userID}
	res, err := client.ProfileClient.GetProfile(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
