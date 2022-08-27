package model

import (
    "github.com/chatpuppy/puppychat/internal/g"
    jsoniter "github.com/json-iterator/go"
    "github.com/liasica/go-encryption/hexutil"
    "github.com/liasica/go-encryption/rsa"
)

type EncryptedKeys interface {
    Encrypt() (hex string, err error)
    Decrypt(hex string) (err error)
}

func KeysEncrypt[T any](pointer *T) (hex string, err error) {
    var data []byte
    data, err = jsoniter.Marshal(pointer)
    if err != nil {
        return
    }
    var b []byte
    b, err = rsa.Encrypt(data, g.RsaPublicKey())
    if err != nil {
        return
    }
    hex = hexutil.Encode(b)
    return
}

func KeysDecrypt[T any](hex string) (pointer *T, err error) {
    pointer = new(T)
    var b []byte
    b, err = hexutil.Decode(hex)
    if err != nil {
        return
    }
    var data []byte
    data, err = rsa.Decrypt(b, g.RsaPrivateKey())
    if err != nil {
        return
    }
    err = jsoniter.Unmarshal(data, pointer)
    return
}
