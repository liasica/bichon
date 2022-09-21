package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/labstack/echo/v4"
)

type message struct{}

var Message = new(message)

// Create
// @ID           MessageCreate
// @Router       /message [POST]
// @Summary      Create message
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        body  body     model.ChatMessage  true  "message detail"
// @Success      200  {object}  app.Response{data=model.ChatMessageCreateRes}  "Response success"
func (*message) Create(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.ChatMessage](c)
    return ctx.SendResponse(service.NewMessage().Create(ctx.Member, req))
}

// List
// @ID           MessageList
// @Router       /message [GET]
// @Summary      List message
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param        query  query   model.MessageListReq  true  "message filter"
// @Success      200  {object}  app.Response{data=[]model.Message}  "Response success"
func (*message) List(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.MessageListReq](c)
    return ctx.SendResponse(service.NewMessage().List(ctx.Member, req))
}
