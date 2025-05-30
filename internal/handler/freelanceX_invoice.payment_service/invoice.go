package handler

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	invoicepb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_invoice.payment_service/invoice"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc/metadata"
	"context"
	"github.com/Prototype-1/freelanceX_apigateway_service/internal/client"
)

func withMetadata(c *gin.Context) context.Context {
	userID := c.GetString("user_id")
	role := c.GetString("role")
	md := metadata.Pairs("user_id", userID, "role", role)
	return metadata.NewOutgoingContext(c.Request.Context(), md)
}

func CreateInvoiceHandler(c *gin.Context) {
	var req struct {
		ClientID       string  `json:"client_id"`
		FreelancerID   string  `json:"freelancer_id"`
		ProjectID      string  `json:"project_id"`
		Type           string  `json:"type"` 
		FixedAmount    float64 `json:"fixed_amount"`
		MilestonePhase string  `json:"milestone_phase"`
		DateFrom       string  `json:"date_from"`
		DateTo         string  `json:"date_to"`  
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateFrom, err := time.Parse(time.RFC3339, req.DateFrom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_from format"})
		return
	}
	dateTo, err := time.Parse(time.RFC3339, req.DateTo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_to format"})
		return
	}

	var invoiceType invoicepb.InvoiceType
	switch req.Type {
	case "FIXED":
		invoiceType = invoicepb.InvoiceType_FIXED
	case "HOURLY":
		invoiceType = invoicepb.InvoiceType_HOURLY
	case "MILESTONE":
		invoiceType = invoicepb.InvoiceType_MILESTONE
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice type"})
		return
	}

	grpcReq := &invoicepb.CreateInvoiceRequest{
		ClientId:       req.ClientID,
		FreelancerId:   req.FreelancerID,
		ProjectId:      req.ProjectID,
		Type:           invoiceType,
		FixedAmount:    req.FixedAmount,
		MilestonePhase: req.MilestonePhase,
		DateFrom:       timestamppb.New(dateFrom),
		DateTo:         timestamppb.New(dateTo),
	}

	resp, err := client.InvoiceClient.CreateInvoice(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp.Invoice)
}

func GetInvoiceHandler(c *gin.Context) {
	invoiceID := c.Param("id")

	grpcReq := &invoicepb.GetInvoiceRequest{
		InvoiceId: invoiceID,
	}

	resp, err := client.InvoiceClient.GetInvoice(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Invoice)
}

func GetInvoicesByUserHandler(c *gin.Context) {
	userID := c.Param("userId")
	role := c.Query("role")

	grpcReq := &invoicepb.GetInvoicesByUserRequest{
		UserId: userID,
		Role:   role,
	}

	resp, err := client.InvoiceClient.GetInvoicesByUser(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Invoices)
}

func GetInvoicesByProjectHandler(c *gin.Context) {
	projectID := c.Param("projectId")

	grpcReq := &invoicepb.GetInvoicesByProjectRequest{
		ProjectId: projectID,
	}

	resp, err :=client. InvoiceClient.GetInvoicesByProject(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Invoices)
}

func UpdateInvoiceStatusHandler(c *gin.Context) {
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoiceID := c.Param("id")

	var status invoicepb.InvoiceStatus
	switch req.Status {
	case "PENDING":
		status = invoicepb.InvoiceStatus_PENDING
	case "PAID":
		status = invoicepb.InvoiceStatus_PAID
	case "CANCELLED":
		status = invoicepb.InvoiceStatus_CANCELLED
	case "OVERDUE":
    status = invoicepb.InvoiceStatus_OVERDUE
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	grpcReq := &invoicepb.UpdateInvoiceStatusRequest{
		InvoiceId: invoiceID,
		Status:    status,
	}

	resp, err := client.InvoiceClient.UpdateInvoiceStatus(withMetadata(c), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.Invoice)
}
