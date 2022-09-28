package model

var (
    // MessageBroadcastChan broadcast to other online members in group
    MessageBroadcastChan chan MessageBroadcast
)

func ChanInitialize() {
    MessageBroadcastChan = make(chan MessageBroadcast, 1024)
}

func SendBroadcast(msg MessageBroadcast) {
    MessageBroadcastChan <- msg
}
