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
// @Success      200  {object}  app.Response{data=model.GroupDetail}  "Response success"
func (*group) Create(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupCreateReq](c)
    return ctx.SendResponse(service.NewGroup().Create(ctx.Member, req))
}
