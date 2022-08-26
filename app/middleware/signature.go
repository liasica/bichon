package middleware

import (
    "fmt"
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/request"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/labstack/echo/v4"
    "io"
    "net/http"
)

var (
    signatureRequired = map[RequestUrlMethod]struct{}{
        "POST/group": {},
    }
)

// Signature signature verifying middleware
func Signature() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // logged member
            var mem *ent.Member
            if ctx, ok := c.(*app.MemberContext); ok && ctx.Member != nil {
                mem = ctx.Member
            }
            rum := RequestUrlMethod(fmt.Sprintf("%s%s", c.Request().Method, c.Path()))
            isPost := c.Request().Method == request.Post
            if _, ok := signatureRequired[rum]; ok && isPost {
                if mem == nil {
                    // need signin
                    return app.Context(c).SendResponse(nil, model.ErrAuthError, http.StatusUnauthorized)
                }

                b, _ := io.ReadAll(c.Request().Body)
                // if invaild signature
                if !request.VerifyingSignature(b, c.Request().Header.Get(app.HeaderSignature), mem.PublicKey) {
                    return app.Context(c).SendResponse(nil, model.ErrInvalidSignature, http.StatusBadRequest)
                }
            }
            return next(c)
        }
    }
}
