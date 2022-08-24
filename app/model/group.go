package model

const (
    // GroupMax Maximum number of groups to create
    GroupMax = 100
)

type GroupCreateReq struct {
    Name  string
    Intro string
}
