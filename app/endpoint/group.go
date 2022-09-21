package endpoint

import (
    "github.com/chatpuppy/puppychat/app"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/service"
    "github.com/labstack/echo/v4"
)

type group struct{}

var Group = new(group)

// Categories
// @ID           GroupCategories
// @Router       /group/categories [GET]
// @Summary      List all group categories
// @Tags         Group
// @Accept       json
// @Produce      json
// @Success      200  {object}  app.Response  "Response success"
func (*group) Categories(c echo.Context) (err error) {
    ctx := app.Context(c)
    return ctx.SendResponse(model.GroupCategories)
}

// Create
// @ID           GroupCreate
// @Router       /group [POST]
// @Summary      Create group
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

// List
// @ID           GroupList
// @Router       /group [GET]
// @Summary      Filter group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        query  query   model.GroupListReq  false  "filter group fields"
// @Success      200  {object}  app.Response{data=model.PaginationRes{items=[]model.GroupListRes}}  "Response success"
func (*group) List(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupListReq](c)
    return ctx.SendResponse(service.NewGroup().List(ctx.Member, req))
}

// JoinedList
// @ID           GroupJoinedList
// @Router       /group/joined [GET]
// @Summary      Joined group list
// @Tags         Group
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.GroupJoinedListRes  "Response success"
func (*group) JoinedList(c echo.Context) (err error) {
    ctx := app.ContextX[app.MemberContext](c)
    return ctx.SendResponse(service.NewGroup().JoinedList(ctx.Member))
}

// Join
// @ID           GroupJoin
// @Router       /group/join [POST]
// @Summary      Join group
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
// @Summary      Share key with group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        body  body     model.GroupShareKeyReq  true  "Group and key info"
// @Success      200  {object}  app.Response{data=model.GroupPublicKey}  "Response success"
func (*group) ShareKey(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupShareKeyReq](c)
    return ctx.SendResponse(service.NewGroup().ShareKey(ctx.Member, req))
}

// Detail
// @ID           GroupDetail
// @Router       /group/{id} [GET]
// @Summary      Detail info of group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Group id"
// @Success      200  {object}  app.Response{data=model.GroupMetaRes}  "Response success"
func (*group) Detail(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupMetaReq](c)
    return ctx.SendResponse(service.NewGroup().Detail(req, ctx.Member))
}

// KeyUsed
// @ID           GroupKeyUsed
// @Router       /group/key/used [POST]
// @Summary      Setting current key of group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        body  body     model.GroupKeyUsedReq  true  "Group keys"
// @Success      200  {object}  app.Response  "Response success"
func (*group) KeyUsed(c echo.Context) (err error) {
    ctx, req := app.MemberContextAndBinding[model.GroupKeyUsedReq](c)
    return ctx.SendResponse(service.NewGroup().KeyUsed(ctx.Member, req))
}
