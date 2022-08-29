package router

import (
    "fmt"
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/middleware"
    "github.com/chatpuppy/puppychat/app/request"
    "github.com/chatpuppy/puppychat/pkg/snag"
    "github.com/labstack/echo/v4"
    mw "github.com/labstack/echo/v4/middleware"
    "github.com/labstack/gommon/log"
    "net/http"
)

func Run() {
    r := echo.New()

    corsConfig := mw.DefaultCORSConfig
    corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, []string{
        app.HeaderContentType,
        app.HeaderMemberAddress,
        app.HeaderSignature,
        app.HeaderTimestamp,
        app.HeaderAuthorization,
    }...)
    corsConfig.ExposeHeaders = append(corsConfig.ExposeHeaders, []string{
        app.HeaderContentType,
        app.HeaderDispositionType,
    }...)

    // load middlewares
    r.Use(
        middleware.Api(),
        middleware.Recover(),
        // logger
        mw.LoggerWithConfig(mw.LoggerConfig{
            Format: `{"time":"${time_custom}","id":"${id}","remote_ip":"${remote_ip}",` +
                `"method":"${method}","uri":"${uri}",` +
                `"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
                `,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
            CustomTimeFormat: "2006-01-02 15:04:05.00000",
        }),
        mw.CORSWithConfig(corsConfig),
        mw.GzipWithConfig(mw.GzipConfig{
            Level: 5,
        }),
        mw.RequestID(),
    )

    r.Validator = request.NewValidator()

    // handle api error
    r.HTTPErrorHandler = func(err error, c echo.Context) {
        ctx := app.Context(c)
        if ctx == nil {
            ctx = app.NewContext(c)
        }
        code := http.StatusBadRequest
        var data any
        switch err.(type) {
        case *snag.Error:
            target := err.(*snag.Error)
            code = target.Code
            data = target.Data
        case *echo.HTTPError:
            target := err.(*echo.HTTPError)
            code = target.Code
            err = fmt.Errorf("%v", target.Message)
        }
        _ = ctx.SendResponse(data, err, code)
    }

    // load app routes
    loadAppRoutes(r)

    // load doc routes
    loadDocRoutes(r)

    log.Fatal(r.Start(":5501"))
}
