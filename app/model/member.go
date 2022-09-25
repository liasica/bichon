package model

import "regexp"

func IsValidAddress(address string) bool {
    return regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(address)
}

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
    Token   string         `json:"token"`
    Profile *MemberProfile `json:"profile"`
}

type MemberProfile struct {
    ID       string `json:"id"`
    Address  string `json:"address"`
    Nickname string `json:"nickname,omitempty"` // optional
    Avatar   string `json:"avatar,omitempty"`   // optional
    Intro    string `json:"intro,omitempty"`    // optional
}

type Member struct {
    ID       string `json:"id"`
    Address  string `json:"address"`
    Nickname string `json:"nickname,omitempty"`
    Avatar   string `json:"avatar,omitempty"`
}

type MemberWithPermission struct {
    Member
    Permission GroupMemberPerm `json:"permission"`
}

type MemberUpdateReq struct {
    Nickname string `json:"nickname"`
    Intro    string `json:"intro"`
}
