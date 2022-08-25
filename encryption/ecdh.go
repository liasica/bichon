package encryption

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "crypto/x509"
    "encoding/hex"
    "fmt"
)

// ECDHGenerateKey generate ecdh private and public keys
func ECDHGenerateKey() (priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey, err error) {
    priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return
    }
    pub = &priv.PublicKey
    return
}

// ECDHPrivateKeyEncode ECDH private key encode to hex string
func ECDHPrivateKeyEncode(priv *ecdsa.PrivateKey) (str string, err error) {
    var b []byte
    b, err = x509.MarshalECPrivateKey(priv)
    if err != nil {
        return
    }
    str = fmt.Sprintf("%x", b)
    return
}

// ECDHPrivateKeyDecode ECDH private key decode from hex string
func ECDHPrivateKeyDecode(str string) (priv *ecdsa.PrivateKey, err error) {
    var b []byte
    b, err = hex.DecodeString(str)
    if err != nil {
        return
    }
    return x509.ParseECPrivateKey(b)
}

// ECDHPublicKeyEncode ECDH public key encode to hex string
func ECDHPublicKeyEncode(pub *ecdsa.PublicKey) string {
    b := elliptic.MarshalCompressed(elliptic.P256(), pub.X, pub.Y)
    return fmt.Sprintf("%x", b)
}

// ECDHPublicKeyDecode ECDH public key decode from hex string
func ECDHPublicKeyDecode(str string) (pub *ecdsa.PublicKey, err error) {
    var b []byte
    b, err = hex.DecodeString(str)
    if err != nil {
        return
    }
    c := elliptic.P256()
    x, y := elliptic.UnmarshalCompressed(c, b)
    pub = &ecdsa.PublicKey{
        Curve: c,
        X:     x,
        Y:     y,
    }
    return
}

// ECDHShareKey exchange ecdh key from compressed public
func ECDHShareKey(compressed string) (shared [sha256.Size]byte, pri *ecdsa.PrivateKey, pub *ecdsa.PublicKey, err error) {
    var data []byte
    data, err = hex.DecodeString(compressed)
    if err != nil {
        return
    }

    c := elliptic.P256()
    x, y := elliptic.UnmarshalCompressed(c, data)
    pua := &ecdsa.PublicKey{
        Curve: c,
        X:     x,
        Y:     y,
    }

    pri, pub, err = ECDHGenerateKey()
    if err != nil {
        return
    }

    a, _ := pua.Curve.ScalarMult(pua.X, pua.Y, pri.D.Bytes())

    shared = sha256.Sum256(a.Bytes())

    return
}
