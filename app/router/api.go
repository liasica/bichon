package router

import (
    "github.com/chatpuppy/puppychat/app/endpoint"
    "github.com/chatpuppy/puppychat/app/middleware"
    "github.com/labstack/echo/v4"
)

func loadAppRoutes(r *echo.Echo) {
    r.Use(
        middleware.Member(),
    )

    // member routes
    r.GET("/member/nonce/:address", endpoint.Member.Nonce)
    r.POST("/member", endpoint.Member.Signin)
    r.GET("/member/:address", endpoint.Member.Profile)

    // group routes
    r.POST("/group", endpoint.Group.Create, middleware.Signature())
}
