package handler

import (
	"os"
	"net/http"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	paymentPb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_invoice.payment_service/payment"
)

type Handler struct {
	PaymentGRPCClient paymentPb.PaymentServiceClient
}

func NewPaymentHandler(client paymentPb.PaymentServiceClient) *Handler {
	return &Handler{PaymentGRPCClient: client}
}

func (h *Handler) SimulatePaymentHandler(c *gin.Context) {
	var req struct {
		InvoiceID   string  `json:"invoice_id" binding:"required"`
		MilestoneID string  `json:"milestone_id"` 
		PayerID     string  `json:"payer_id" binding:"required"`
		ReceiverID  string  `json:"receiver_id" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := c.GetString("role")
	if role != "client" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only clients can initiate payments"})
		return
	}

	grpcReq := &paymentPb.SimulatePaymentRequest{
		InvoiceId:   req.InvoiceID,
		MilestoneId: req.MilestoneID,
		PayerId:     req.PayerID,
		ReceiverId:  req.ReceiverID,
		Amount:      req.Amount,
	}

	md := metadata.New(map[string]string{
		"role": role,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := h.PaymentGRPCClient.SimulatePayment(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment_id":      res.PaymentId,
		"amount_paid":     res.AmountPaid,
		"platform_fee":    res.PlatformFee,
		"amount_credited": res.AmountCredited,
		"status":          res.Status,
		 "razorpay_order_id": res.RazorpayOrderId, 
    "razorpay_key_id":  os.Getenv("RAZORPAY_KEY_ID"),
	})
}
