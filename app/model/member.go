package model

type MemberAddressParam struct {
    Address string `json:"address" param:"address" validate:"required,address"`
}

type MemberNonceRes struct {
    Nonce string `json:"nonce"`
}

type MemberSigninReq struct {
    Address   string `json:"address" validate:"required,address" trans:"User's public ethereum address'"`
    Signature string `json:"signature" validate:"required"`
    Nonce     string `json:"nonce" validate:"required"`
}

type MemberSigninRes struct {
    Token string
}
