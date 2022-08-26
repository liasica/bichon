package router

import (
    "github.com/chatpuppy/puppychat/app/endpoint"
    "github.com/chatpuppy/puppychat/app/middleware"
    "github.com/labstack/echo/v4"
)

func loadAppRoutes(r *echo.Echo) {
    mr := r.Group("/member")
    mr.GET("/nonce/:address", endpoint.Member.Nonce)
    mr.POST("", endpoint.Member.Signin)
    mr.GET("/:address", endpoint.Member.Profile)

    r.Use(
        middleware.Member(),
    )
}
