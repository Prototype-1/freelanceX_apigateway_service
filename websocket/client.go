package websocket

import (
 "github.com/gorilla/websocket"
    "log"
)

type Client struct {
    Conn   *websocket.Conn
    UserID string
    Send   chan MessagePayload
    Hub    *Hub
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
