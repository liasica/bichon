package service

import (
    "context"
    "crypto/rsa"
    "entgo.io/ent"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/group"
    "github.com/chatpuppy/puppychat/internal/ent/groupmember"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/internal/ent/member"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    rsaTools "github.com/liasica/go-encryption/rsa"
    log "github.com/sirupsen/logrus"
)

type syncService struct {
    ctx context.Context
}

func NewSync() *syncService {
    return &syncService{
        ctx: context.Background(),
    }
}

func (s *syncService) SaveSyncData(req *model.SyncData, priv *rsa.PrivateKey) {
    // saving data
    data, err := rsaTools.DecryptUsePrivateKey(req.Value, priv)
    if err != nil {
        log.Errorf("[D] reading sync data failed: %v", err)
        return
    }

    switch req.Table {
    case group.Table:
        err = NewGroup().SaveSyncData(data, ent.Op(req.Op))
    case groupmember.Table:
        err = NewGroupMember().SaveSyncData(data, ent.Op(req.Op))
    case key.Table:
        err = NewKey().SaveSyncData(data, ent.Op(req.Op))
    case member.Table:
        err = NewMember().SaveSyncData(data, ent.Op(req.Op))
    case message.Table:
        err = NewMessage().SaveSyncData(data, ent.Op(req.Op))
    }
    if err != nil {
        log.Errorf("[D] sync %s data failed: %v", req.Table, err)
    }
}
