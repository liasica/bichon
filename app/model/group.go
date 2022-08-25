package model

const (
    // GroupMax Maximum number of groups to create
    GroupMax = 100
)

type GroupCreateReq struct {
    Name  string
    Intro string
}

type GroupJoinReq struct {
    GroupID uint64 `json:"groupId" validate:"required" trans:"Group ID"`
}
