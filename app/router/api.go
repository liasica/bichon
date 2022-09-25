package router

import (
    "github.com/chatpuppy/puppychat/app/endpoint"
    "github.com/chatpuppy/puppychat/app/middleware"
    "github.com/chatpuppy/puppychat/app/ws"
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
    g.POST("/member/update", endpoint.Member.Update, middleware.Signature())

    // group routes
    g.POST("/group", endpoint.Group.Create, middleware.Signature())
    g.GET("/group/joined", endpoint.Group.JoinedList)
    g.GET("/group/categories", endpoint.Group.Categories)
    g.GET("/group", endpoint.Group.List)
    g.POST("/group/join", endpoint.Group.Join, middleware.Signature())
    g.POST("/group/leave", endpoint.Group.Leave, middleware.Signature())
    g.POST("/group/key", endpoint.Group.ShareKey)
    g.GET("/group/:id", endpoint.Group.Detail)
    g.POST("/group/key/used", endpoint.Group.KeyUsed)
    g.POST("/group/key/update", endpoint.Group.Update, middleware.Signature())
    g.POST("/group/invite/recode", endpoint.Group.ReCode, middleware.Signature())
    g.POST("/group/active", endpoint.Group.Active)
    g.POST("/group/kickout", endpoint.Group.Kickout, middleware.Signature())
    g.POST("/group/manager", endpoint.Group.Manager, middleware.Signature())

    // message
    // websocket
    g.GET("/message/:address", func(c echo.Context) error {
        return ws.Serve(ws.Hub, c)
    })
    g.POST("/message", endpoint.Message.Create)
    g.GET("/message", endpoint.Message.List)
    g.POST("/message/read", endpoint.Message.Read)
}
