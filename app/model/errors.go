package model

import (
    "errors"
    "net/http"
)

var (
    ErrAuthError         = errors.New("auth failed")
    ErrInvalidAddress    = errors.New("invalid address")
    ErrInvalidSignature  = errors.New("invalid signature")
    ErrMaximunGroups     = errors.New("created group has exceeded the maximum number")
    ErrMaximunMembers    = errors.New("group's members has exceeded the maximum number")
    ErrCreateFrequency   = errors.New("create group frequency limited")
    ErrAlreadyInGroup    = errors.New("already in the group")
    ErrNotFoundGroup     = errors.New("not found group")
    ErrKeyShareFrequency = errors.New("key generate frequency limited")
)

func ErrStatus(err error) int {
    switch err {
    case ErrNotFoundGroup:
        return http.StatusNotFound
    default:
        return http.StatusBadRequest
    }
}
