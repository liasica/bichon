package encryption

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
)

// RSAEncrypt RSA encrypt to []byte (using public key)
func RSAEncrypt(data, key []byte) (b []byte, err error) {
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
    return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// RSAEncryptToBase64 RSA encrypt to base64 string (using public key)
func RSAEncryptToBase64(data, key []byte) (b64 string, err error) {
    var b []byte
    b, err = RSAEncrypt(data, key)
    if err != nil {
        return
    }
    b64 = base64.StdEncoding.EncodeToString(b)
    return
}

// RSADecrypt RSA decode bytes to bytes (using private key)
func RSADecrypt(b, key []byte) (data []byte, err error) {
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

    // decode data
    data, err = rsa.DecryptPKCS1v15(rand.Reader, priv, b)
    if err != nil {
        return
    }
    return
}

// RSADecryptFromBase64 RSA decode base64 string to bytes (using private key)
func RSADecryptFromBase64(b64 string, key []byte) (data []byte, err error) {
    var b []byte
    b, err = base64.StdEncoding.DecodeString(b64)
    if err != nil {
        return
    }

    // decode data
    data, err = RSADecrypt(b, key)
    if err != nil {
        return
    }
    return
}
