package service

import (
    "context"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
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

func (s *groupMemberService) Query(memID, groupID string) (*ent.GroupMember, error) {
    return s.orm.Query().Where(groupmember.GroupID(groupID), groupmember.MemberID(memID)).First(s.ctx)
}
