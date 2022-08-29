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
