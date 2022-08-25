package encryption

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
)

// RSAEncrypt RSA encrypt to base64 string (using public key)
func RSAEncrypt(data []byte, pubKey string) (b64 string, err error) {
    var key []byte
    key, err = base64.StdEncoding.DecodeString(pubKey)
    if err != nil {
        return
    }
    block, _ := pem.Decode(key)
    if block == nil {
        err = errors.New("RSA Key error")
        return
    }
    var dpub any
    dpub, err = x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return
    }
    pub, ok := dpub.(*rsa.PublicKey)
    if !ok {
        err = errors.New("RSA Key parse error")
        return
    }
    var b []byte
    b, err = rsa.EncryptPKCS1v15(rand.Reader, pub, data)
    if err != nil {
        return
    }
    b64 = base64.StdEncoding.EncodeToString(b)

    return
}

// RSADecrypt RSA decode base64 string to bytes (using private key)
func RSADecrypt(b64 string, privkey string) (data []byte, err error) {
    var key []byte
    key, err = base64.StdEncoding.DecodeString(privkey)
    if err != nil {
        return
    }
    block, _ := pem.Decode(key)
    if block == nil {
        err = errors.New("RSA Key error")
        return
    }

    // parse private
    var priv *rsa.PrivateKey
    priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return
    }

    var b []byte
    b, err = base64.StdEncoding.DecodeString(b64)
    if err != nil {
        return
    }

    // decode data
    data, err = rsa.DecryptPKCS1v15(rand.Reader, priv, b)
    if err != nil {
        return
    }
    return
}
