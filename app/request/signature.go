package request

import (
    "bytes"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

// VerifyingSignature verify the authenticiy of the signature
func VerifyingSignature(data []byte, signature string, pub string) (matches bool) {
    publicKeyBytes, err := hexutil.Decode(pub)
    if err != nil {
        return
    }
    hash := crypto.Keccak256Hash(data)
    var sigPublicKey []byte
    sigPublicKey, err = crypto.Ecrecover(hash.Bytes(), []byte(signature))
    if err != nil {
        return
    }
    matches = bytes.Equal(sigPublicKey, publicKeyBytes)
    return
}
