package websocket

type Hub struct {
    clients    map[string]*Client        
    register   chan *Client
    unregister chan *Client
    broadcast  chan MessagePayload
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[string]*Client),
        register:   make(chan *Client),
        unregister: make(chan *Client),
        broadcast:  make(chan MessagePayload),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client.UserID] = client
        case client := <-h.unregister:
            if _, ok := h.clients[client.UserID]; ok {
                delete(h.clients, client.UserID)
                close(client.Send)
            }
        case msg := <-h.broadcast:
            if recipient, ok := h.clients[msg.ToUserID]; ok {
                recipient.Send <- msg
            }
        }
    }
}
