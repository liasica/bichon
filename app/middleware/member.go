package middleware

import (
    "fmt"
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/labstack/echo/v4"
    "net/http"
)

var (
    memberTokenPrefix = "Bearer "
    memberAuthSkipper = map[RequestUrlMethod]struct{}{
        "GET/member/nonce/:address": {},
        "POST/member":               {},
        "GET/member/:address":       {},
        "GET/message/:address":      {},
        "GET/group/categories":      {},
        "GET/group":                 {},
    }
)

func Member() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) (err error) {
            ctx := app.NewMemberContext(c.(*app.BaseContext), nil)

            rum := RequestUrlMethod(fmt.Sprintf("%s%s", c.Request().Method, c.Path()))
            _, ok := memberAuthSkipper[rum]
            needLogin := !ok

            h := c.Request().Header
            address := h.Get(app.HeaderMemberAddress)
            prefix := len(memberTokenPrefix)
            token := h.Get(app.HeaderAuthorization)

            var mem *ent.Member
            if model.IsValidAddress(address) && len(token) > len(memberTokenPrefix) {
                token = token[prefix:]
                mem, _ = service.NewMember().Authenticate(address, token)
            }

            // if need auth but member is nil
            if needLogin && mem == nil {
                return ctx.SendResponse(nil, model.ErrAuthRequired, http.StatusUnauthorized)
            }

            ctx.Member = mem

            return next(ctx)
        }
    }
}
