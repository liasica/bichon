package distribution

import (
    "bufio"
    "context"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/tea"
    "github.com/go-redis/redis/v8"
    jsoniter "github.com/json-iterator/go"
    rsaTools "github.com/liasica/go-encryption/rsa"
    log "github.com/sirupsen/logrus"
    "net"
    "time"
)

const (
    syncedIndexKey = "SYNCED_INDEX"
)

type client struct {
    net.Conn
}

// get synced index
func (c *client) syncedIndex() (index uint64) {
    index, _ = cache.Get(context.Background(), syncedIndexKey).Uint64()
    return
}

// store synced index
func (c *client) storeSyncedIndex(index uint64) {
    // meaning next sync index is + 1
    cache.Set(context.Background(), syncedIndexKey, index+1, redis.KeepTTL)
}

func CreateClient() {
    addr := g.Config.App.Distribution.String()
    log.Infof("[D] connecting distribution node: %s", addr)

    conn, err := net.Dial("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }

    c := &client{
        conn,
    }

    go c.run()
}

func (c *client) run() {
    defer func(conn net.Conn) {
        _ = conn.Close()
    }(c.Conn)

    c.startSync()

    go c.readBump()
    go c.broadcast()

    wait := make(chan struct{})

    <-wait
}

// send message to dribution node
func (c *client) sendRequest(msg *model.SyncRequest) (err error) {
    msg.NodeID = g.NodeID()
    msg.Nonce = []byte(fmt.Sprintf("%d", time.Now().UnixNano()))

    hasher := sha256.New()
    hasher.Write(msg.Nonce)
    sum := hasher.Sum(nil)

    // signature nonce string
    msg.Signature, err = rsa.SignPSS(rand.Reader, g.RsaPrivateKey(), crypto.SHA256, sum, nil)
    if err != nil {
        return
    }

    _, err = c.Write(msg.Marshal())
    return
}

func (c *client) startSync() {
    index := c.syncedIndex()

    msg := &model.SyncRequest{SyncedStart: tea.UInt64(index), ApiUrl: tea.String(g.ApiUrl())}

    err := c.sendRequest(msg)
    if err != nil {
        log.Fatal(err)
    }
}

func (c *client) readBump() {
    reader := bufio.NewReader(c)

    for {
        b, _, err := reader.ReadLine()
        if err != nil {
            log.Fatalf("[D] client read failed: %v", err)
        }

        c.readResponse(b)
    }
}

func (c *client) readResponse(b []byte) {
    fmt.Println("readResponse", string(b))
    var res model.SyncResponse
    err := jsoniter.Unmarshal(b, &res)
    if err != nil {
        log.Errorf("[D] read response failed: %v, respon raw data: %s", err, b)
        return
    }

    if res.Error != nil {
        log.Fatalf("[D] receive error message: %v", *res.Error)
    }

    if res.Progress != nil {
        log.Infof("[D] node sync progress: %.2f%%", *res.Progress)

        if *res.Progress == 100 {
            model.DistributionNodeSyncedChan <- true
        }
    }

    if res.Data == nil {
        return
    }

    // saving sync data
    if res.Data.Value != nil {
        service.NewSync().SaveSyncData(res.Data, g.RsaPrivateKey())
    }

    // updating index
    c.storeSyncedIndex(res.Data.Index)
}

// send modify data to distribution node
func (c *client) broadcast() {
    for {
        select {
        // broadcast sync data
        case data := <-model.DistributionBroadcastChan:
            go c.writeSyncData(data)
        }
    }
}

func (c *client) writeSyncData(data *model.SyncData) {
    var err error
    data.Value, err = rsaTools.EncryptUsePublicKey(data.Value, g.RsaPublicKey())
    if err != nil {
        log.Errorf("[D] sync data encrypt failed: %v", err)
    }
    msg := &model.SyncRequest{Data: data}
    err = c.sendRequest(msg)
    if err != nil {
        log.Errorf("[D] send sync request failed: %v", err)
    }
}
