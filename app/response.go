package app

import (
    "bytes"
    jsoniter "github.com/json-iterator/go"
    "net/http"
)

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
}

// SendResponse Send api response
func (c *BaseContext) SendResponse(data any, err error, params ...int) error {
    var r *Response
    if err != nil {
        r = &Response{
            Message: err.Error(),
            Code:    http.StatusBadRequest,
        }
    } else {
        r = &Response{
            Message: "ok",
            Data:    data,
            Code:    http.StatusOK,
        }
    }

    if len(params) > 0 {
        r.Code = params[0]
    }

    buffer := &bytes.Buffer{}
    encoder := jsoniter.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    _ = encoder.Encode(r)

    return c.JSONBlob(http.StatusOK, buffer.Bytes())
}
