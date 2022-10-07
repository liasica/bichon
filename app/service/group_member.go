package service

import (
    "context"
    "github.com/chatpuppy/puppychat/internal/ent"
)

type groupMemberService struct {
    ctx context.Context
    orm *ent.GroupMemberClient
}

func NewGroupMember() *groupMemberService {
    return &groupMemberService{
        ctx: context.Background(),
        orm: ent.Database.GroupMember,
    }
}

func (s *groupMemberService) SaveSyncData(b []byte, op ent.Op) (err error) {
    return ent.SaveGroupMemberSyncData(b, op, nil)
}
