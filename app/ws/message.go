package ws

import (
    jsoniter "github.com/json-iterator/go"
)

type Operate int

const (
    OperateRegister Operate = iota + 1 // register operate
    OperateChat
)

const (
    OperateOK = "OK"
)

type Message struct {
    Operate Operate `json:"operate"`
    Data    any     `json:"data"`
}

func (m *Message) Marshal() ([]byte, error) {
    return jsoniter.Marshal(m)
}

func NewMessage[T any](op Operate, data T) *Message {
    return &Message{
        Operate: op,
        Data:    data,
    }
}
