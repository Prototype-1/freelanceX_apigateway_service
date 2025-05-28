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

func (h *Handler) CreatePaymentOrderHandler(c *gin.Context) {
	var req struct {
		InvoiceID   string  `json:"invoice_id" binding:"required"`
		MilestoneID string  `json:"milestone_id"` // Optional
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

	grpcReq := &paymentPb.CreatePaymentOrderRequest{
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

	res, err := h.PaymentGRPCClient.CreatePaymentOrder(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment_id":        res.PaymentId,
		"razorpay_order_id": res.RazorpayOrderId,
		"amount":            res.Amount,
		"currency":          res.Currency,
		"invoice_id":        res.InvoiceId,
		"razorpay_key_id":   os.Getenv("RAZORPAY_KEY_ID"),
	})
}

func (h *Handler) RazorpayCheckoutPageHandler(c *gin.Context) {
	orderID := c.Query("order_id")
	amount := c.Query("amount")
	currency := c.Query("currency")
	invoiceID := c.Query("invoice_id")
	razorpayKey := os.Getenv("RAZORPAY_KEY_ID")

	if orderID == "" || amount == "" || currency == "" || invoiceID == "" {
		c.String(http.StatusBadRequest, "Missing required query parameters")
		return
	}

	htmlContent := `
<!DOCTYPE html>
<html>
<head>
	<title>Razorpay Checkout</title>
	<script src="https://checkout.razorpay.com/v1/checkout.js"></script>
</head>
<body>
	<h1>Complete your Payment</h1>
	<button id="rzp-button">Pay Now</button>

	<script>
		var options = {
			"key": "` + razorpayKey + `",
			"amount": ` + amount + `, 
			"currency": "` + currency + `",
			"order_id": "` + orderID + `",
			"handler": function (response){
				fetch("/api/payment/verify", {
					method: "POST",
					headers: {
						"Content-Type": "application/json"
					},
					body: JSON.stringify({
						razorpay_payment_id: response.razorpay_payment_id,
						razorpay_order_id: response.razorpay_order_id,
						razorpay_signature: response.razorpay_signature,
						invoice_id: "` + invoiceID + `"
					})
				}).then(res => res.json()).then(data => {
					alert(data.message);
				}).catch(err => {
					alert("Verification failed, please contact support");
				});
			},
			"prefill": {
				"name": "",
				"email": "",
				"contact": ""
			},
			"theme": {
				"color": "#3399cc"
			}
		};
		var rzp1 = new Razorpay(options);
		document.getElementById('rzp-button').onclick = function(e){
			rzp1.open();
			e.preventDefault();
		}
	</script>
</body>
</html>
`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}

func (h *Handler) VerifyPaymentHandler(c *gin.Context) {
    var req struct {
        RazorpayPaymentId string `json:"razorpay_payment_id" binding:"required"`
        RazorpayOrderId   string `json:"razorpay_order_id" binding:"required"`
        RazorpaySignature string `json:"razorpay_signature" binding:"required"`
        InvoiceID         string `json:"invoice_id" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    grpcReq := &paymentPb.VerifyPaymentRequest{
        RazorpayPaymentId: req.RazorpayPaymentId,
        RazorpayOrderId:   req.RazorpayOrderId,
        RazorpaySignature: req.RazorpaySignature,
        InvoiceId:         req.InvoiceID,
    }

    res, err := h.PaymentGRPCClient.VerifyPayment(c.Request.Context(), grpcReq)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if !res.Valid {
        c.JSON(http.StatusBadRequest, gin.H{"message": res.Message})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": res.Message})
}
