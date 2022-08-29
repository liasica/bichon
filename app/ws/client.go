package ws

import (
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/gorilla/websocket"
)

type Client struct {
    mem  *ent.Member
    hub  *Hub
    conn *websocket.Conn
    send chan []byte
}

// Disconnect close member's client
// same member's client can not connect at the same time
func (c *Client) Disconnect() {
    _ = c.conn.Close()
}

func NewClient(hub *Hub, mem *ent.Member, conn *websocket.Conn) *Client {
    return &Client{
        mem:  mem,
        hub:  hub,
        conn: conn,
        // TODO use larger buffer or chunks message
        send: make(chan []byte, 256),
    }
}
