package g

import (
    "crypto/rsa"
    "encoding/base64"
    "github.com/bwmarrin/snowflake"
    jsoniter "github.com/json-iterator/go"
    rsaTools "github.com/liasica/go-encryption/rsa"
)

type Node struct {
    PublicKey  []byte `json:"publicKey"`
    PrivateKey []byte `json:"privateKey"`
    NodeID     int64  `json:"nodeId"`
    ApiUrl     string `json:"apiUrl"`

    RsaPublicKey  *rsa.PublicKey  `json:"-"`
    RsaPrivateKey *rsa.PrivateKey `json:"-"`

    snowflaker *snowflake.Node
}

var node *Node

func IinitializeNode(str string) (err error) {
    node = new(Node)
    node.snowflaker, err = snowflake.NewNode(node.NodeID)
    if err != nil {
        return
    }
    return node.Unmarshal(str)
}

func (n *Node) Marshal() (string, error) {
    b, err := jsoniter.Marshal(n)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(b), nil
}

func (n *Node) Unmarshal(str string) (err error) {
    var b []byte
    b, err = base64.StdEncoding.DecodeString(str)
    if err != nil {
        return
    }

    err = jsoniter.Unmarshal(b, n)
    if err != nil {
        return
    }

    // parse keys
    n.RsaPrivateKey, err = rsaTools.ParsePrivateKey(n.PrivateKey)
    if err != nil {
        return
    }

    n.RsaPublicKey, err = rsaTools.ParsePublicKey(n.PublicKey)
    return
}

func RsaPublicKey() *rsa.PublicKey {
    return node.RsaPublicKey
}

func RsaPrivateKey() *rsa.PrivateKey {
    return node.RsaPrivateKey
}

func IsDistributionNode() bool {
    return node.NodeID == 0
}

func IsDistributionNodeID(id int64) bool {
    return id == 0
}

func NodeID() int64 {
    return node.NodeID
}

func ApiUrl() string {
    return node.ApiUrl
}

func Snowflaker() *snowflake.Node {
    return node.snowflaker
}
