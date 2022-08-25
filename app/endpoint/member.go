package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/labstack/echo/v4"
)

type member struct{}

var Member = new(member)

// Nonce
// @ID           MemberNonce
// @Router       /member/nonce/{address} [GET]
// @Summary      Get Member Signin Nonce String
// @Tags         Member
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response{data=model.MemberNonceRes}  "Response success"
func (*member) Nonce(c echo.Context) (err error) {
    ctx, req := app.ContextBinding[model.MemberAddressParam](c)
    return ctx.SendResponse(service.NewMember().Nonce(req))
}

// Signin
// @ID           MemberSignin
// @Router       /memeber [POST]
// @Summary      Member Signin
// @Tags         Member
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response{data=model.MemberSigninRes}  "Response success"
func (*member) Signin(c echo.Context) (err error) {
    ctx, req := app.ContextBinding[model.MemberSigninReq](c)
    return ctx.SendResponse(service.NewMember().Signin(req))
}
