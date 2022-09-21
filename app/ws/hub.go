package ws

import (
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "sync"
)

var Hub *hub

type hub struct {
    // online clients current
    // type is menberID => *Client
    clients sync.Map

    // register requests from the clients
    register chan *Client

    // unregister requests from clients
    unregister chan *Client
}

func CreateHub() {
    Hub = &hub{
        register:   make(chan *Client, 256),
        unregister: make(chan *Client, 256),
        clients:    sync.Map{},
    }
    go Hub.Run()
}

func (h *hub) Run() {
    for {
        select {
        case client := <-h.register:
            // check if member registered now
            if exists, ok := h.clients.Load(client.mem.ID); ok {
                exists.(*Client).Disconnect()
            }
            // store client
            h.clients.Store(client.mem.ID, client)
        case client := <-h.unregister:
            // unregister client
            client.Disconnect()
        case message := <-service.BroadcastChan:
            go h.Broadcast(message)
        }
    }
}

// Broadcast broadcast message to other online members
func (h *hub) Broadcast(req *model.ChatMessage) {
    // get members
    memIDs := cache.Group.Members(req.GroupID)
    for _, memID := range memIDs {
        if exists, ok := h.clients.Load(memID); ok {
            // send message to client
            exists.(*Client).send <- &Message{
                Operate: OperateChat,
                Data:    req,
            }
        }
    }
}
