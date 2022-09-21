package ent

import (
    "context"
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    log "github.com/sirupsen/logrus"
)

// Cache cache group's member
func (m *GroupMemberMutation) Cache(delete bool) {
    memberID, _ := m.MemberID()
    groupID, _ := m.GroupID()

    if memberID == "" || groupID == "" {
        return
    }

    log.Infof("[Group Cache] cache groupID: %s, memberID: %s", groupID, memberID)

    // check group cache if initialized
    if cache.Group.Exists(memberID) {
        if delete {
            cache.Group.Remove(groupID, memberID)
            return
        }
    } else {
        memberIDs, _ := Database.GroupMember.Query().
            Where(groupmember.GroupID(groupID), groupmember.MemberIDNotIn(memberID)).
            Select(groupmember.FieldMemberID).
            Strings(context.Background())
        cache.Group.Initialize(groupID, memberIDs)
    }

    cache.Group.Add(groupID, memberID)
}
