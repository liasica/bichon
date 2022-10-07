package model

var (
    // MessageBroadcastChan broadcast to other online members in group
    MessageBroadcastChan chan MessageBroadcast

    // DistributionBroadcastChan broadcast sync data channel
    DistributionBroadcastChan chan *SyncData

    // DistributionFromClientChan sync data from clients
    DistributionFromClientChan chan *SyncData

    // DistributionNodeSyncedChan if client synced
    DistributionNodeSyncedChan chan bool
)

func ChanInitialize() {
    MessageBroadcastChan = make(chan MessageBroadcast, 1024)
    DistributionBroadcastChan = make(chan *SyncData)
    DistributionFromClientChan = make(chan *SyncData)
    DistributionNodeSyncedChan = make(chan bool)
}

func SendBroadcast(msg MessageBroadcast) {
    MessageBroadcastChan <- msg
}
