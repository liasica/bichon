package router

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/gommon/log"
)

func Run() {
    r := echo.New()

    loadDocRoutes(r)

    log.Fatal(r.Start(":5501"))
}
