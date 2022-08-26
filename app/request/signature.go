package request

import (
    "bytes"
    "crypto/ecdsa"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/ethereum/go-ethereum/accounts"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

// VerifyingSignature verify the authenticiy of the signature
func VerifyingSignature(data []byte, signature string, mem *ent.Member) (matches bool) {
    var err error
    data = []byte(utils.Md5String(data))
    sig := hexutil.MustDecode(signature)
    sig[crypto.RecoveryIDOffset] -= 27
    msg := accounts.TextHash(data)
    var recovered *ecdsa.PublicKey
    recovered, err = crypto.SigToPub(msg, sig)
    if err != nil {
        return
    }

    // getting member's public key
    var publicKeyBytes []byte
    publicKeyBytes, err = hexutil.Decode(mem.PublicKey)
    if err != nil {
        return
    }

    // is signature public key equals to member's public key
    matches = bytes.Equal(crypto.FromECDSAPub(recovered), publicKeyBytes)

    // verify signature
    signatureNoRecoverID := sig[:len(sig)-1]
    verified := crypto.VerifySignature(publicKeyBytes, msg, signatureNoRecoverID)

    matches = matches && verified
    return
}
