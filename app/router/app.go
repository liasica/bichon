package router

import (
    "github.com/chatpuppy/puppychat/app/endpoint"
    "github.com/labstack/echo/v4"
)

func loadAppRoutes(r *echo.Echo) {
    mr := r.Group("/member")
    mr.GET("/:address", endpoint.Member.Nonce)
    mr.POST("", endpoint.Member.Signin)
}
