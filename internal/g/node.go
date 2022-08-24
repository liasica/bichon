package g

import (
    "encoding/base64"
    jsoniter "github.com/json-iterator/go"
)

type Node struct {
    PublicKey  []byte `json:"publicKey"`
    PrivateKey []byte `json:"privateKey"`
    NodeID     uint16 `json:"nodeId"`
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

func NodeID() uint16 {
    return node.NodeID
}
