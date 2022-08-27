package app

import (
    "bytes"
    "github.com/chatpuppy/puppychat/app/model"
    jsoniter "github.com/json-iterator/go"
    "github.com/modern-go/reflect2"
    "net/http"
)

const (
    CodeEmpty = 0
)

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
}

// SendResponse Send api response
func (c *BaseContext) SendResponse(data any, params ...any) error {
    r := &Response{
        Message: "ok",
    }
    for _, param := range params {
        switch p := param.(type) {
        case error:
            r.Message = p.Error()
            if r.Code == CodeEmpty {
                r.Code = model.ErrStatus(p)
            }
            break
        case int:
            r.Code = p
            break
        }
    }

    if r.Code == CodeEmpty {
        r.Code = http.StatusOK
    }

    if data != nil {
        r.Data = data
    }

    if reflect2.IsNil(data) {
        r.Data = nil
    }

    buffer := &bytes.Buffer{}
    encoder := jsoniter.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    _ = encoder.Encode(r)

    return c.JSONBlob(http.StatusOK, buffer.Bytes())
}
