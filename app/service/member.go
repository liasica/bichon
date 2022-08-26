package service

import (
    "context"
    "crypto/ecdsa"
    "crypto/rand"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/ethereum/go-ethereum/accounts"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/golang-jwt/jwt/v4"
    "math/big"
    "strings"
    "time"
)

type memberService struct {
    ctx context.Context
    orm *ent.MemberClient
}

func NewMember() *memberService {
    return &memberService{
        ctx: context.Background(),
        orm: ent.Database.Member,
    }
}

// QueryAddress getting member from address
func (s *memberService) QueryAddress(address string) (*ent.Member, error) {
    return s.orm.Query().Where(member.Address(address)).First(s.ctx)
}

// Authenticate getting and Authenticate member
func (s *memberService) Authenticate(address string, token string) (*ent.Member, error) {
    mem, err := s.QueryAddress(address)
    if err != nil {
        return nil, err
    }
    jws := NewJwt(mem.Nonce, mem.Address, 24*time.Hour)
    var jrc *jwt.RegisteredClaims
    jrc, err = jws.Verify(token)
    if err != nil {
        return nil, err
    }
    if jrc.Subject == mem.Address {
        return mem, nil
    }
    return nil, model.ErrAuthError
}

// Nonce generate auth nonce string for member's address
func (s *memberService) Nonce(req *model.MemberAddressParam) (res *model.MemberNonceRes, err error) {
    res = new(model.MemberNonceRes)
    max := new(big.Int)
    max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
    n, _ := rand.Int(rand.Reader, max)
    res.Nonce = n.Text(16)

    // turn address to all lowercase characters
    address := strings.ToLower(req.Address)
    // find member, if not exists, create it
    mem, _ := s.orm.Query().Where(member.Address(address)).First(s.ctx)
    if mem == nil {
        _, err = s.orm.Create().SetAddress(address).SetNonce(res.Nonce).Save(s.ctx)
        if err != nil {
            return
        }
    } else {
        _, err = mem.Update().SetNonce(res.Nonce).Save(s.ctx)
        if err != nil {
            return
        }
    }
    return
}

// Signin member signin, return jwt token
func (s *memberService) Signin(req *model.MemberSigninReq) (res *model.MemberSigninRes, err error) {
    // getting member from database with nonce string
    mem, _ := s.orm.Query().Where(member.Address(req.Address), member.Nonce(req.Nonce)).First(s.ctx)
    if mem == nil {
        err = model.ErrAuthError
        return
    }
    // decode the provided signature
    sig := hexutil.MustDecode(req.Signature)
    // see at https://github.com/ethereum/go-ethereum/blob/master/internal/ethapi/api.go#L516
    sig[crypto.RecoveryIDOffset] -= 27
    // hash nonce
    msg := accounts.TextHash([]byte(req.Nonce))
    // recover the public key that signed data
    var pub *ecdsa.PublicKey
    pub, err = crypto.SigToPub(msg, sig)
    if err != nil {
        return
    }
    // create an ethereum address from the extracted public key
    recoveredAddr := crypto.PubkeyToAddress(*pub)
    // validate member's address
    if req.Address != strings.ToLower(recoveredAddr.Hex()) {
        err = model.ErrAuthError
        return
    }
    // generate jwt token
    var token string
    token, err = NewJwt(req.Nonce, req.Address, 24*time.Hour).CreateStandard(req.Address)
    if err != nil {
        return
    }
    res = &model.MemberSigninRes{Token: token, Profile: s.Profile(mem)}
    return
}

// Profile returns member's profile
func (s *memberService) Profile(param any) *model.MemberProfile {
    var mem *ent.Member
    switch p := param.(type) {
    case string:
        mem, _ = s.QueryAddress(p)
        break
    case *ent.Member:
        mem = p
        break
    default:
        return nil
    }
    if mem == nil {
        return nil
    }
    return &model.MemberProfile{
        Address:  mem.Address,
        Nickname: mem.Nickname,
        Avatar:   mem.Avatar,
        Intro:    mem.Intro,
    }
}
