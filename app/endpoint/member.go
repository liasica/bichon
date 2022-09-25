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
// @Summary      Member signin nonce
// @Description  Getting nonce string to be signed
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
// @Summary      Member signin
// @Tags         Member
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response{data=model.MemberSigninRes}  "Response success"
func (*member) Signin(c echo.Context) (err error) {
    ctx, req := app.ContextBinding[model.MemberSigninReq](c)
    return ctx.SendResponse(service.NewMember().Signin(req))
}

// Profile
// @ID           MemberProfile
// @Router       /member/{address} [GET]
// @Summary      Member profile
// @Description  Getting member's profile using address
// @Tags         Member
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.MemberProfile  "Response success"
func (*member) Profile(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.MemberAddressParam](c)
    return ctx.SendResponse(service.NewMember().Profile(req.Address))
}

// Update
// @ID           MemberUpdate
// @Router       /member/update [POST]
// @Summary      Update member's info
// @Description  need signature verify
// @Tags         Member
// @Accept       json
// @Produce      json
// @Param        body  body     model.MemberUpdateReq  true  "Member's info"
// @Success      200  {object}  app.Response  "Response success"
func (*member) Update(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.MemberUpdateReq](c)
    return ctx.SendResponse(nil, service.NewMember().Update(ctx.Member, req))
}
