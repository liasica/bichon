package g

import (
    "encoding/base64"
    "github.com/bwmarrin/snowflake"
    jsoniter "github.com/json-iterator/go"
    log "github.com/sirupsen/logrus"
)

type Node struct {
    PublicKey  []byte `json:"publicKey"`
    PrivateKey []byte `json:"privateKey"`
    NodeID     int64  `json:"nodeId"`
}

var node *Node

func IinitializeNode(str string) error {
    node = new(Node)
    return node.Unmarshal(str)
}

func (n *Node) Marshal() (string, error) {
    b, err := jsoniter.Marshal(n)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(b), nil
}

func (n *Node) Unmarshal(str string) error {
    b, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return err
    }
    err = jsoniter.Unmarshal(b, n)
    return err
}

func RsaPublicKey() []byte {
    return node.PublicKey
}

func RsaPrivateKey() []byte {
    return node.PrivateKey
}

func NodeID() int64 {
    return node.NodeID
}

func SnowflakeNode() *snowflake.Node {
    n, err := snowflake.NewNode(node.NodeID)
    if err != nil {
        log.Fatal(err)
    }
    return n
}
