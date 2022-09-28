package ws

import (
    "fmt"
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/app/model"
    log "github.com/sirupsen/logrus"
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
    // channel init
    model.ChanInitialize()

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
        case message := <-model.MessageBroadcastChan:
            go h.broadcast(message)
        }
    }
}

// broadcast message to other online members
func (h *hub) broadcast(req model.MessageBroadcast) {
    // get members
    memIDs := cache.Group.Members(req.GroupID)
    str := fmt.Sprintf("%s has %d online members, ", req.GroupAddress, len(memIDs))
    n := 0
    for _, memID := range memIDs {
        if exists, ok := h.clients.Load(memID); ok {
            c := exists.(*Client)
            // TODO skip owner
            if req.Member.ID == c.mem.ID {
                // continue
            }
            // send message to client
            c.send <- &Message{
                Operate: OperateChat,
                Data:    &req.Message,
            }
            n += 1
        }
    }
    str += fmt.Sprintf("sent %d members", n)
    log.Info(str)
}
