package router

import (
    "github.com/chatpuppy/puppychat/app/endpoint"
    "github.com/chatpuppy/puppychat/app/middleware"
    "github.com/labstack/echo/v4"
)

func loadAppRoutes(r *echo.Echo) {
    g := r.Group("")
    g.Use(
        middleware.Member(),
    )

    // member routes
    g.GET("/member/nonce/:address", endpoint.Member.Nonce)
    g.POST("/member", endpoint.Member.Signin)
    g.GET("/member/:address", endpoint.Member.Profile)

    // group routes
    g.POST("/group", endpoint.Group.Create, middleware.Signature())
    g.POST("/group/join", endpoint.Group.Join, middleware.Signature())
    g.POST("/group/key", endpoint.Group.ShareKey)
}
