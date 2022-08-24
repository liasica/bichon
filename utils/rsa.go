package utils

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
)

type rsaUtil struct {
    bits int
}

func NewRsa() *rsaUtil {
    return &rsaUtil{
        bits: 2048,
    }
}

// GenRsaKey Generate rsa keys
func (u *rsaUtil) GenRsaKey() (private, public []byte) {
    // Generate privite key
    privateKey, err := rsa.GenerateKey(rand.Reader, u.bits)
    if err != nil {
        panic(err)
    }
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }
    private = pem.EncodeToMemory(block)

    // Generate public key
    publicKey := &privateKey.PublicKey
    var derPkix []byte
    derPkix, err = x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        panic(err)
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    public = pem.EncodeToMemory(block)
    return
}
