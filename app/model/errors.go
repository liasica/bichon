package model

import "errors"

var (
    ErrAuthError        = errors.New("auth failed")
    ErrInvalidAddress   = errors.New("invalid address")
    ErrInvalidSignature = errors.New("invalid signature")
)
