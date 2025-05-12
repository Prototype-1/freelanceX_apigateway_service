package websocket

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // TODO: tighten in prod
    },
}

func ServeWS(hub *Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Assume AuthMiddleware already ran
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
            Conn:   conn,
            UserID: userID,
            Send:   make(chan MessagePayload),
            Hub:    hub,
        }

        hub.register <- client

        go client.WritePump()
        go client.ReadPump()
    }
}
