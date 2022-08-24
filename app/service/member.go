package service

import (
    "context"
    "github.com/chatpuppy/puppychat/internal/ent"
)

type memberService struct {
    ctx context.Context
    orm *ent.MemberClient
}

func NewMember() *memberService {
    return &memberService{
        ctx: context.Background(),
        orm: ent.Database.Member,
    }
}
