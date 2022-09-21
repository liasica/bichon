package ws

import (
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/gorilla/websocket"
    jsoniter "github.com/json-iterator/go"
    log "github.com/sirupsen/logrus"
    "time"
)

const (
    // Time allowed to write a message to the peer.
    writeWait = 10 * time.Second

    // Time allowed to read the next pong message from the peer.
    pongWait = 60 * time.Second

    // Send pings to peer with this period. Must be less than pongWait.
    pingPeriod = (pongWait * 9) / 10

    // Maximum message size allowed from peer.
    maxMessageSize = 15 << 10 << 10
)

type Client struct {
    mem     *ent.Member
    hub     *hub
    conn    *websocket.Conn
    address string
    send    chan *Message
}

func safeCloseCh(ch chan *Message) (closed bool) {
    defer func() {
        if recover() != nil {
            closed = true
        }
    }()

    close(ch)
    closed = true
    return
}

// Disconnect close member's client
// same member's client can not connect at the same time
func (c *Client) Disconnect() {
    _ = c.conn.Close()
    if c.mem != nil {
        c.hub.clients.Delete(c.mem.ID)
    }
    safeCloseCh(c.send)
}

func NewClient(hub *hub, conn *websocket.Conn, address string) *Client {
    return &Client{
        hub:     hub,
        conn:    conn,
        address: address,
        send:    make(chan *Message, 256),
    }
}

func (c *Client) writePump() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        c.hub.unregister <- c
    }()

    for {
        select {
        case message, ok := <-c.send:
            _ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if !ok {
                // The hub closed the channel.
                _ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            go c.writeMessage(message)
        case <-ticker.C:
            _ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}

func (c *Client) writeMessage(message *Message) {
    w, err := c.conn.NextWriter(websocket.BinaryMessage)
    if err != nil {
        log.Errorf("[WS] message write failed: %s", err)
        return
    }

    if message.Operate == OperateChat {
        cm, ok := message.Data.(*model.ChatMessage)
        if !ok {
            return
        }
        // get group key
        var keys *model.GroupMemberRawKeys
        keys, err = model.LoadSharedKeys(c.mem.ID, cm.GroupID)
        if err != nil {
            return
        }
        err = cm.Encrypt(keys)
        if err != nil {
            log.Errorf("[WS] message data encrypt error: %v", err)
            return
        }
    }

    b, _ := message.Marshal()

    _, err = w.Write(b)
    if err != nil {
        log.Errorf("[WS] message write failed: %s", err)
        return
    }

    if err = w.Close(); err != nil {
        return
    }
}

func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
    }()

    c.conn.SetReadLimit(maxMessageSize)
    _ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
    c.conn.SetPongHandler(func(string) error {
        _ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })
    for {
        t, r, err := c.conn.NextReader()
        if err != nil {
            if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Errorf("[WS] error: %v", err)
            }
            // TODO error send
            return
        }

        switch t {
        case websocket.BinaryMessage:
            msg := new(Message)
            err = jsoniter.NewDecoder(r).Decode(msg)
            if err != nil {
                // TODO error send
                log.Errorf("[WS] read message error: %v", err)
                break
            }

            switch msg.Operate {
            case OperateRegister:
                go c.doRegister(msg)
            }
        }
    }
}

// do client register operate
func (c *Client) doRegister(msg *Message) {
    var err error
    defer func() {
        data := OperateOK
        if err != nil {
            data = err.Error()
        }
        c.send <- NewMessage(OperateRegister, data)
    }()

    token, ok := msg.Data.(string)
    if !ok {
        err = model.ErrAuthRequired
        return
    }
    // read token and find member
    var mem *ent.Member
    mem, err = service.NewMember().Authenticate(c.address, token)
    if err != nil {
        return
    }

    // register client
    c.mem = mem
    c.hub.register <- c
    return
}
