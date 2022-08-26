package app

import (
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/labstack/echo/v4"
)

type MemberContext struct {
    *BaseContext
    Member *ent.Member
}

func NewMemberContext(ctx *BaseContext, mem *ent.Member) *MemberContext {
    return &MemberContext{
        BaseContext: ctx,
        Member:      mem,
    }
}

func MemberContextAndBinding[T any](c echo.Context) (*MemberContext, *T) {
    return ContextBindingX[MemberContext, T](c)
}
