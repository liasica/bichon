package model

import "errors"

var (
    ErrAuthError        = errors.New("auth failed")
    ErrInvalidAddress   = errors.New("invalid address")
    ErrInvalidSignature = errors.New("invalid signature")
    ErrMaximunGroups    = errors.New("created group has exceeded the maximum number")
)
