package app

import "github.com/labstack/echo/v4"

type BaseContext struct {
    echo.Context
}

func Context(c echo.Context) *BaseContext {
    return c.(*BaseContext)
}
