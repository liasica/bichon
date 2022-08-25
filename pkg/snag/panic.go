package snag

import (
    "fmt"
    "net/http"
)

func Panic(params ...any) {
    panic(NewError(params...))
}

func NewError(params ...any) *Error {
    out := &Error{
        Code: http.StatusBadRequest,
    }

    for _, param := range params {
        switch param.(type) {
        case string:
            out.Message = param.(string)
            break
        case error:
            out.Message = param.(error).Error()
            break
        case int:
            out.Code = param.(int)
            break
        default:
            out.Data = param
            break
        }
    }

    return out
}

func WithPanicError(cb func()) (err error) {
    defer func() {
        if v := recover(); v != nil {
            err = fmt.Errorf("%v", v)
        }
    }()

    cb()

    return
}
