package websocket

import (
    "context"
 "github.com/gorilla/websocket"
    "log"
    "encoding/json"
    "time"
    "github.com/Prototype-1/freelanceX_apigateway_service/kafka"
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

    if c.MessageClient == nil {
        log.Println("MessageClient is nil in ReadPump")
    }
    if c.Hub == nil {
        log.Println("Hub is nil in ReadPump")
    }
    if c.Conn == nil {
        log.Println("Conn is nil in ReadPump")
    }

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
          kafkaEvent := map[string]interface{}{
            "from_user_id": msg.FromUserID,
            "to_user_id":   msg.ToUserID,
            "project_id":   msg.ProjectID,
            "content":      msg.Message,
            "timestamp":    time.Now().Format(time.RFC3339),
        }
        jsonEvent, err := json.Marshal(kafkaEvent)
        if err == nil {
            _ = kafka.SendMessage("new.message", msg.FromUserID, jsonEvent)
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
