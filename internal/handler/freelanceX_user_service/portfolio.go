package handler

import (
	"context"
	"net/http"
	"time"
	pb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/portfolio"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/handler/freelanceX_user_service/utils"
)

func CreatePortfolio(c *gin.Context) {
	var req struct {
		FreelancerID string `json:"freelancer_id" binding:"required"`
		Title        string `json:"title" binding:"required"`
		Description  string `json:"description"`
		ImageURL     string `json:"image_url"`
		Link         string `json:"link"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := utils.InjectMetadataFromGin(c)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := client.PortfolioClient.CreatePortfolio(ctx, &pb.CreatePortfolioRequest{
		FreelancerId: req.FreelancerID,
		Title:        req.Title,
		Description:  req.Description,
		ImageUrl:     req.ImageURL,
		Link:         req.Link,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Portfolio created"})
}

func GetPortfolio(c *gin.Context) {
	freelancerID := c.Param("freelancer_id")

	ctx := utils.InjectMetadataFromGin(c)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := client.PortfolioClient.GetPortfolio(ctx, &pb.GetPortfolioRequest{
		FreelancerId: freelancerID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeletePortfolio(c *gin.Context) {
	portfolioID := c.Param("portfolio_id")

	ctx := utils.InjectMetadataFromGin(c)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()


	_, err := client.PortfolioClient.DeletePortfolio(ctx, &pb.DeletePortfolioRequest{
		PortfolioId: portfolioID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted"})
}
