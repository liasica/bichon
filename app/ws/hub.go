package ws

import (
    "sync"
)

type Hub struct {
    // online clients current
    clients map[string]*Client

    // register requests from the clients
    register chan *Client

    // unregister requests from clients
    unregister chan *Client

    // cached used group member keys, value is *model.GroupMemberRawKeys
    keys sync.Map
}

func NewHub() *Hub {
    return &Hub{
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            // check if member registered now
            if c, ok := h.clients[client.mem.ID]; ok {
                go c.Disconnect()
            }
            // store client
            h.clients[client.mem.ID] = client
        case client := <-h.unregister:
            if c, ok := h.clients[client.mem.ID]; ok {
                c.Disconnect()
                delete(h.clients, client.mem.ID)
                close(client.send)
            }
        }
    }
}
