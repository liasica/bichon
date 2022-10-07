package distribution

import (
    "crypto/rsa"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/fsnotify/fsnotify"
    log "github.com/sirupsen/logrus"
    "io/fs"
    "os"
    "path/filepath"
    "regexp"
)

const (
    nodesDirectory = "config/nodes"
)

var (
    re = regexp.MustCompile(`.*?\d+.node$`)
)

type Node struct {
    NodeID     int64
    ApiUrl     string
    PrivateKey *rsa.PrivateKey
    PublicKey  *rsa.PublicKey
    Synced     bool
}

func (h *hub) watchNodes() {
    if _, err := os.Stat(nodesDirectory); os.IsNotExist(err) {
        _ = os.MkdirAll(nodesDirectory, 0755)
    }

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer func(watcher *fsnotify.Watcher) {
        _ = watcher.Close()
    }(watcher)

    h.getNodes()

    go func() {
        for {
            select {
            case ev, ok := <-watcher.Events:
                if !ok {
                    log.Fatal("[D] file watcher event error")
                    return
                }
                if len(re.FindStringIndex(ev.Name)) > 0 && (ev.Op != fsnotify.Chmod) {
                    h.getNodes()
                }
            case err, _ = <-watcher.Errors:
                log.Fatal("[D] error:", err)
            }
        }
    }()

    err = watcher.Add(nodesDirectory)
    if err != nil {
        log.Fatalf("[D] watcher error: %v", err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}

func (h *hub) getNodes() {
    _ = filepath.Walk(nodesDirectory, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return nil
        }

        if info.IsDir() {
            return nil
        }

        if len(re.FindStringIndex(path)) < 0 {
            return nil
        }

        var b []byte
        b, err = os.ReadFile(path)
        if err != nil {
            return nil
        }

        n := new(g.Node)
        err = n.Unmarshal(string(b))
        if err != nil {
            return nil
        }

        node := &Node{
            NodeID:     n.NodeID,
            ApiUrl:     n.ApiUrl,
            PrivateKey: n.RsaPrivateKey,
            PublicKey:  n.RsaPublicKey,
        }

        conn, ok := h.clients.Load(node)
        if !ok {
            conn = nil
        }
        h.clients.Store(node, conn)

        return nil
    })
}

// OnlineNodes list all online nodes
func (h *hub) OnlineNodes() (ns []string) {
    ns = []string{
        g.ApiUrl(),
    }

    h.clients.Range(func(key, value any) bool {
        if value != nil {
            ns = append(ns, key.(*Node).ApiUrl)
        }
        return true
    })

    return ns
}

// getNode get node from node id
func (h *hub) getNode(id int64) (node *Node, err error) {

    h.clients.Range(func(key, _ any) bool {
        if key.(*Node).NodeID == id {
            node = key.(*Node)
            return false
        }
        return true
    })

    if node == nil {
        err = model.ErrNotFoundNode
    }

    return
}
