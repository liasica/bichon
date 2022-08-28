package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/labstack/echo/v4"
)

type group struct{}

var Group = new(group)

// Create
// @ID           GroupCreate
// @Router       /group [POST]
// @Summary      Group Create
// @Description  create an group (need signature verify)
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        body  body     model.GroupCreateReq  true  "Group info"
// @Success      200  {object}  app.Response{data=model.GroupDetailWithPublicKey}  "Response success"
func (*group) Create(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupCreateReq](c)
    return ctx.SendResponse(service.NewGroup().Create(ctx.Member, req))
}

// Join
// @ID           GroupJoin
// @Router       /group/join [POST]
// @Summary      Join Group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        body  body     model.GroupJoinReq  true  "Join info"
// @Success      200  {object}  app.Response{data=model.GroupDetailWithPublicKey}  "Response success"
func (*group) Join(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupJoinReq](c)
    return ctx.SendResponse(service.NewGroup().Join(ctx.Member, req))
}

// ShareKey
// @ID           GroupShareKey
// @Router       /group/key [POST]
// @Summary      Share key with Group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        body  body     model.GroupShareKeyReq  true  "Group and key info"
// @Success      200  {object}  app.Response{data=model.GroupPublicKey}  "Response success"
func (*group) ShareKey(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupShareKeyReq](c)
    return ctx.SendResponse(service.NewGroup().ShareKey(ctx.Member, req))
}
