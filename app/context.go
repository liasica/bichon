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

// Context getting base context
func Context(c echo.Context) *BaseContext {
    switch ctx := c.(type) {
    case *BaseContext:
        return ctx
    case *MemberContext:
        return ctx.BaseContext
    default:
        return nil
    }
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

type ContextWrapper interface {
    MemberContext
}

func ContextX[T ContextWrapper](c any) *T {
    ctx, _ := c.(*T)
    return ctx
}

func ContextBindingX[K ContextWrapper, T any](c echo.Context) (*K, *T) {
    ctx := ContextX[K](c)
    req := new(T)
    Context(c).BindValidate(req)
    return ctx, req
}
