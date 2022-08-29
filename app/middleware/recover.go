package middleware

import (
    "fmt"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/snag"
    "github.com/labstack/echo/v4"
    log "github.com/sirupsen/logrus"
    "runtime/debug"
)

func Recover() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            defer func() {
                if r := recover(); r != nil {
                    stack := string(debug.Stack())
                    var message string
                    sp := true
                    switch t := r.(type) {
                    case *snag.Error:
                        c.Error(t)
                        message = t.Error()
                        sp = false
                    case *ent.ValidationError:
                        c.Error(t.Unwrap())
                        message = t.Error()
                    case error:
                        message = t.Error()
                        c.Error(t)
                    default:
                        message = fmt.Sprintf("%v", r)
                        c.Error(fmt.Errorf("%v", r))
                    }

                    if sp && g.Config.App.PrintStack {
                        log.Error(fmt.Sprintf("%s\n%s", message, stack))
                    }
                }
            }()
            return next(c)
        }
    }
}
