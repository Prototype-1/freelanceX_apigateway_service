package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	reviewPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_user_service/review"
)

type SubmitReviewRequest struct {
	ProjectID    string `json:"project_id" binding:"required"`
	FreelancerID string `json:"freelancer_id" binding:"required"`
	Rating       int32  `json:"rating" binding:"required,min=1,max=5"`
	Feedback     string `json:"feedback"`
}

func SubmitReview(c *gin.Context) {
	var req SubmitReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	res, err := client.ReviewClient.SubmitReview(c, &reviewPb.ReviewRequest{
		ProjectId:    req.ProjectID,
		FreelancerId: req.FreelancerID,
		ClientId:     clientID.(string),
		Rating:       req.Rating,
		Feedback:     req.Feedback,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"review_id":    res.Id,
		"project_id":   res.ProjectId,
		"freelancer_id": res.FreelancerId,
		"client_id":    res.ClientId,
		"rating":       res.Rating,
		"feedback":     res.Feedback,
		"created_at":   res.CreatedAt,
	})
}

func GetFreelancerReviews(c *gin.Context) {
	freelancerID := c.Param("freelancer_id")
	if freelancerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "freelancer_id is required"})
		return
	}

	res, err := client.ReviewClient.GetFreelancerReviews(c, &reviewPb.GetReviewRequest{FreelancerId: freelancerID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var reviews []gin.H
	for _, r := range res.Reviews {
		parsedTime, _ := time.Parse(time.RFC3339, r.CreatedAt)
		reviews = append(reviews, gin.H{
			"id":            r.Id,
			"project_id":    r.ProjectId,
			"freelancer_id": r.FreelancerId,
			"client_id":     r.ClientId,
			"rating":        r.Rating,
			"feedback":      r.Feedback,
			"created_at":    parsedTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
