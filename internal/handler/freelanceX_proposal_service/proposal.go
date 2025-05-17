package handler

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
	pb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_proposal_service"
	"fmt"
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc/metadata"
)

func withMetadata(c *gin.Context) context.Context {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	md := metadata.Pairs("user_id", userID, "role", role)
	return metadata.NewOutgoingContext(c.Request.Context(), md)
}

func CreateProposalHandler(c *gin.Context) {
	var req struct {
		ClientID     string `json:"client_id"`
		FreelancerID string `json:"freelancer_id"`
		TemplateID   string `json:"template_id"`
		Title        string `json:"title"`
		Content      string `json:"content"`
		Status       string `json:"status"` 
		Version      int32  `json:"version"`
		Deadline     string `json:"deadline"` 
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.CreateProposalRequest{
		ClientId:     req.ClientID,
		FreelancerId: req.FreelancerID,
		TemplateId:   req.TemplateID,
		Title:        req.Title,
		Content:      req.Content,
		Status:       req.Status,
		DeadlineStr:  req.Deadline,
	}

	resp, err := client.ProposalClient.CreateProposal(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetProposalByIDHandler(c *gin.Context) {
	proposalID := c.Param("id")

	resp, err := client.ProposalClient.GetProposalByID(withMetadata(c), &pb.GetProposalRequest{
		ProposalId: proposalID,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateProposalHandler(c *gin.Context) {
	proposalID := c.Param("id")

	var req struct {
		Title       string `json:"title"`
		Content     string `json:"content"`
		Version     int32  `json:"version"`
		Deadline    string `json:"deadline"` 
		Status   string `json:"status"` 
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.UpdateProposalRequest{
		ProposalId:  proposalID,
		Title:       req.Title,
		Content:     req.Content,
		Version:     req.Version,
		DeadlineStr: req.Deadline,
		Status:     req.Status,
	}

	if req.Deadline != "" {
	if t, err := time.Parse(time.RFC3339, req.Deadline); err == nil {
		grpcReq.Deadline = timestamppb.New(t)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deadline format"})
		return
	}
}

	resp, err := client.ProposalClient.UpdateProposal(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ListProposalsHandler(c *gin.Context) {
	clientID := c.Query("client_id")
	freelancerID := c.Query("freelancer_id")
	status := c.Query("status")
	skip := c.DefaultQuery("skip", "0")
	limit := c.DefaultQuery("limit", "10")

	skipVal, err1 := parseInt64(skip)
	limitVal, err2 := parseInt64(limit)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid skip or limit"})
		return
	}

	grpcReq := &pb.ListProposalsRequest{
		ClientId:     clientID,
		FreelancerId: freelancerID,
		Status:       status,
		Skip:         skipVal,
		Limit:        limitVal,
	}

	resp, err := client.ProposalClient.ListProposals(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func SaveTemplateHandler(c *gin.Context) {
	var req struct {
		FreelancerID string `json:"freelancer_id"`
		Title        string `json:"title"`
		Content      string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.SaveTemplateRequest{
		FreelancerId: req.FreelancerID,
		Title:        req.Title,
		Content:      req.Content,
	}

	resp, err := client.ProposalClient.SaveTemplate(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetTemplatesHandler(c *gin.Context) {
	freelancerID := c.Param("freelancer_id")

	resp, err := client.ProposalClient.GetTemplatesForFreelancer(withMetadata(c), &pb.GetTemplatesRequest{
		FreelancerId: freelancerID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func parseInt64(s string) (int64, error) {
	var i int64
	_, err := fmt.Sscan(s, &i)
	return i, err
}
