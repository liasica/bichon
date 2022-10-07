package distribution

import (
    "context"
    "crypto"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/tea"
    jsoniter "github.com/json-iterator/go"
    rsaTools "github.com/liasica/go-encryption/rsa"
    "github.com/panjf2000/gnet/v2"
    log "github.com/sirupsen/logrus"
    "math"
    "os"
    "os/signal"
    "sort"
    "sync"
    "syscall"
)

const (
    distributionKey = "DISTRIBUTION"
)

var (
    Hub *hub
)

type hub struct {
    gnet.BuiltinEventEngine

    // address to listen
    addr string

    // current online nodes
    // the key is *Node, and value is gnet.Conn
    clients sync.Map
}

func (h *hub) saveSyncedData(data *model.SyncData) {

    ctx := context.Background()
    data.Index, _ = cache.LLen(ctx, distributionKey).Uint64()

    cache.RPush(ctx, distributionKey, data)
}

func (h *hub) dataNotSynced(start uint64) (items []*model.SyncData) {
    ctx := context.Background()
    stop, _ := cache.LLen(ctx, distributionKey).Uint64()
    if stop < start {
        return
    }
    err := cache.LRange(ctx, distributionKey, int64(start), int64(stop)).ScanSlice(&items)
    if err != nil {
        log.Errorf("[D] getting not synced data fail: %v", err)
    }

    sort.Slice(items, func(i, j int) bool {
        return items[i].Index < items[j].Index
    })

    return
}

func (h *hub) OnBoot(eng gnet.Engine) gnet.Action {
    log.Infof("distribution node server started at %s", h.addr)
    return gnet.None
}

func (h *hub) OnTraffic(c gnet.Conn) (action gnet.Action) {
    action = gnet.None

    buf, err := c.Next(-1)
    if err != nil {
        return
    }

    b := buf[:len(buf)-1]

    go h.readRequest(c, b)

    return gnet.None
}

func (h *hub) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
    return
}

func (h *hub) OnClose(c gnet.Conn, err error) (action gnet.Action) {
    log.Infof("[D] client closed: fd = %d, error?: %v", c.Fd(), err)
    h.disconnect(c)
    return
}

func (h *hub) Close(c gnet.Conn, err error) {
    _, _ = c.Write(model.SyncResError(err))

    _ = c.Close()

    id := h.disconnect(c)

    log.Infof("[D] close connect fd = %d, nodeid = %d", c.Fd(), id)
}

func (h *hub) disconnect(c gnet.Conn) (id int64) {

    h.clients.Range(func(key, value any) bool {
        if value == nil {
            return true
        }
        conn := value.(gnet.Conn)
        if conn.Fd() == c.Fd() {
            node := key.(*Node)
            id = node.NodeID
            // str += fmt.Sprintf(", nodeid = %d", key.(*Node).NodeID)
            h.clients.Store(key, nil)
            return false
        }
        return true
    })

    return
}

func CreateHub() {
    Hub = &hub{
        addr: fmt.Sprintf(":%d", g.Config.App.Distribution.Port),
    }

    go Hub.run()
    go Hub.signal()
    go Hub.broadcast()
    go Hub.watchNodes()
}

func (h *hub) run() {
    addr := fmt.Sprintf("tcp://%s", h.addr)
    log.Fatal(gnet.Run(h, addr, gnet.WithMulticore(true)))
}

func (h *hub) signal() {
    c := make(chan os.Signal)
    signal.Notify(c)
    for s := range c {
        switch s {
        case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Kill, os.Interrupt, syscall.SIGKILL:
            _ = gnet.Stop(context.Background(), fmt.Sprintf("tcp://%s", h.addr))
        }
    }
}

func (h *hub) broadcast() {
    for {
        select {
        // broadcast sync data
        case data := <-model.DistributionBroadcastChan:
            // saving data
            h.saveSyncedData(data)

            // broadcast data
            h.clients.Range(func(key, value any) bool {
                node := key.(*Node)
                if c, ok := value.(gnet.Conn); ok && node.Synced {
                    if data.NodeID == node.NodeID {
                        data.Value = nil
                    }
                    go h.sendSyncResToNode(data, node, c, nil, nil)
                }
                return true
            })
        }
    }
}

func (h *hub) sendSyncResToNode(data *model.SyncData, node *Node, c gnet.Conn, current *uint64, progress *float64) {
    var err error
    if data != nil && data.Value != nil {
        data.Value, err = rsaTools.EncryptUsePublicKey(data.Value, node.PublicKey)
        if err != nil {
            log.Errorf("[D] encrypt data value failed: %v", err)
            return
        }
    }
    res := &model.SyncResponse{Data: data, Current: current, Progress: progress}
    _, err = c.Write(res.Marshal())
    if err != nil {
        log.Errorf("[D] send sync data to node (nodeid = %d) failed: %v", node.NodeID, err)
    }
    return
}

func (h *hub) readRequest(c gnet.Conn, b []byte) {
    // getting data
    var req model.SyncRequest
    err := jsoniter.Unmarshal(b, &req)
    if err != nil {
        // close client
        h.Close(c, err)
        return
    }

    // getting node
    var node *Node
    node, err = h.getNode(req.NodeID)
    if err != nil {
        h.Close(c, err)
        return
    }

    // verify nonce signature
    hash := sha256.New()
    _, _ = hash.Write(req.Nonce)
    sum := hash.Sum(nil)
    err = rsa.VerifyPSS(node.PublicKey, crypto.SHA256, sum, req.Signature, nil)
    if err != nil {
        log.Errorf("[D] nodeid = %d, %v", node.NodeID, model.ErrSignature)
        h.Close(c, model.ErrSignature)
        return
    }

    // register message
    // store connection
    if req.Data == nil {
        node.ApiUrl = *req.ApiUrl
        h.clients.Store(node, c)
    }

    // sync start request
    if req.SyncedStart != nil {
        go h.startSync(*req.SyncedStart, node, c)
        return
    }

    if req.Data != nil {
        // saving sync data
        go service.NewSync().SaveSyncData(req.Data, node.PrivateKey)
    }
}

func (h *hub) startSync(start uint64, node *Node, c gnet.Conn) {
    items := h.dataNotSynced(start)
    total := float64(len(items))
    if total > 0 {
        current := start
        var progress float64 = 0
        for i, item := range items {
            current = uint64(i) + start
            progress = math.Round(float64(100*(i+1))/total*100.00) / 100
            h.sendSyncResToNode(item, node, c, tea.UInt64(current), tea.Float64(progress))
        }
    } else {
        h.sendSyncResToNode(nil, node, c, tea.UInt64(1), tea.Float64(100))
    }
    // update node synced
    node.Synced = true
    h.clients.Store(node, c)
}
