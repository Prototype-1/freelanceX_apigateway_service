package websocket

import (
    "context"
 "github.com/gorilla/websocket"
    "log"
   pb  "github.com/Prototype-1/freelanceX_apigateway_service/proto/freelanceX_message.notification_service"
)

type Client struct {
    Conn          *websocket.Conn
    UserID        string
    Send          chan MessagePayload
    Hub           *Hub
    MessageClient pb.MessageServiceClient 
    Ctx           context.Context
}

type MessagePayload struct {
    FromUserID string `json:"from_user_id"`
    ToUserID   string `json:"to_user_id"`
    ProjectID  string `json:"project_id"`
    Message    string `json:"message"`
}

func (c *Client) ReadPump() {
    defer func() {
        c.Hub.unregister <- c
        c.Conn.Close()
    }()

    for {
        var msg MessagePayload
        if err := c.Conn.ReadJSON(&msg); err != nil {
            log.Println("read error:", err)
            break
        }

        msg.FromUserID = c.UserID
        _, err := c.MessageClient.SendMessage(c.Ctx, &pb.SendMessageRequest{
            FromUserId: msg.FromUserID,
            ToUserId:   msg.ToUserID,
            ProjectId:  msg.ProjectID,
            Message:    msg.Message,
            Attachments: []string{}, 
        })
        if err != nil {
            log.Printf("gRPC send failed: %v", err)
            continue
        }
        c.Hub.broadcast <- msg
    }
}


func (c *Client) WritePump() {
    for msg := range c.Send {
        if err := c.Conn.WriteJSON(msg); err != nil {
            log.Println("write error:", err)
            break
        }
    }
}
