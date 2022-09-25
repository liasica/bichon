package model

var (
    // BroadcastChan broadcast to other online members in group
    BroadcastChan chan Message
)

func ChanInitialize() {
    BroadcastChan = make(chan Message, 1024)
}

func SendBroadcast(msg Message) {
    BroadcastChan <- msg
}
