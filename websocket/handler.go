package websocket

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
    "context"
    pb "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_message.notification_service"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true 
    },
}


func ServeWS(hub *Hub, messageClient pb.MessageServiceClient) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        if userID == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }

        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            return
        }

        client := &Client{
            Conn:          conn,
            UserID:        userID,
            Send:          make(chan MessagePayload),
            Hub:           hub,
            MessageClient: messageClient,
            Ctx:           context.Background(), 
        }

        hub.register <- client

        go client.WritePump()
        go client.ReadPump()
    }
}
