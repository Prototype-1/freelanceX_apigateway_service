package handler

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	milestonePb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_invoice.payment_service/milestone"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateMilestoneRuleHandler(c *gin.Context) {
	var req struct {
		ProjectID string  `json:"project_id"`
		Phase     string  `json:"phase"`
		Amount    float64 `json:"amount"`
		DueDate   string  `json:"due_date"` 
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := uuid.Parse(req.ProjectID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project_id"})
		return
	}

	grpcReq := &milestonePb.CreateMilestoneRuleRequest{
		ProjectId: req.ProjectID,
		Phase:     req.Phase,
		Amount: req.Amount,
	}

	if req.DueDate != "" {
		dueTime, err := time.Parse(time.RFC3339, req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format"})
			return
		}
		grpcReq.DueDate = timestamppb.New(dueTime)
	}

	resp, err := client.MilestoneClient.CreateMilestoneRule(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func UpdateMilestoneRuleHandler(c *gin.Context) {
	var req struct {
		ID      string  `json:"id"`
		Phase   string  `json:"phase"`
		Amount  float64 `json:"amount"`
		DueDate string  `json:"due_date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := uuid.Parse(req.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid milestone rule ID"})
		return
	}

	grpcReq := &milestonePb.UpdateMilestoneRuleRequest{
		Id:     req.ID,
		Phase:  req.Phase,
		Amount: req.Amount,
	}

	if req.DueDate != "" {
		dueTime, err := time.Parse(time.RFC3339, req.DueDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format"})
			return
		}
		grpcReq.DueDate = timestamppb.New(dueTime)
	}

	resp, err := client.MilestoneClient.UpdateMilestoneRule(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetMilestonesByProjectIDHandler(c *gin.Context) {
	projectID := c.Param("project_id")

	if _, err := uuid.Parse(projectID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project_id"})
		return
	}

	grpcReq := &milestonePb.GetMilestonesByProjectIDRequest{
		ProjectId: projectID,
	}

	resp, err := client.MilestoneClient.GetMilestonesByProjectID(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetMilestoneByProjectIDAndPhaseHandler(c *gin.Context) {
	projectID := c.Param("project_id")
	phase := c.Param("phase")

	if _, err := uuid.Parse(projectID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid project_id"})
		return
	}

	grpcReq := &milestonePb.GetMilestoneByProjectIDAndPhaseRequest{
		ProjectId: projectID,
		Phase:     phase,
	}

	resp, err := client.MilestoneClient.GetMilestoneByProjectIDAndPhase(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
