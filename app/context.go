package app

import (
    "github.com/labstack/echo/v4"
)

type BaseContext struct {
    echo.Context
}

// NewContext create context
func NewContext(c echo.Context) *BaseContext {
    return &BaseContext{
        Context: c,
    }
}

func Context(c echo.Context) *BaseContext {
    ctx, ok := c.(*BaseContext)
    if ok {
        return ctx
    }
    return nil
}

// BindValidate Bind and validate data
func (c *BaseContext) BindValidate(ptr any) {
    err := c.Bind(ptr)
    if err != nil {
        panic(err)
    }
    err = c.Validate(ptr)
    if err != nil {
        panic(err)
    }
}

// ContextBinding bind data and context
func ContextBinding[T any](c echo.Context) (*BaseContext, *T) {
    ctx := Context(c)
    req := new(T)
    ctx.BindValidate(req)
    return ctx, req
}
