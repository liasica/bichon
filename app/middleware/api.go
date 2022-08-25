package middleware

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/labstack/echo/v4"
)

func Api() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            return next(app.NewContext(c))
        }
    }
}
