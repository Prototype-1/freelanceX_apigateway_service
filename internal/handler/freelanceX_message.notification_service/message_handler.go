package handlers

import (
    "context"
    "net/http"
    "time"
	"strconv"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    pb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_message.notification_service"
)

type MessageHandler struct {
    Client pb.MessageServiceClient
}

func NewMessageHandler(client pb.MessageServiceClient) *MessageHandler {
    return &MessageHandler{Client: client}
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
    senderID := c.Query("sender_id")
    receiverID := c.Query("receiver_id")
    limit := c.DefaultQuery("limit", "20")
    offset := c.DefaultQuery("offset", "0")

    if _, err := uuid.Parse(senderID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sender_id"})
        return
    }
    if _, err := uuid.Parse(receiverID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receiver_id"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req := &pb.GetMessagesRequest{
        SenderId:   senderID,
        ReceiverId: receiverID,
        Limit:      int32(toInt(limit, 20)),
        Offset:     int32(toInt(offset, 0)),
    }

    resp, err := h.Client.GetMessages(ctx, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"messages": resp.Messages})
}

func toInt(s string, defaultVal int) int {
    if val, err := strconv.Atoi(s); err == nil {
        return val
    }
    return defaultVal
}
