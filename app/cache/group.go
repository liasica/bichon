package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
)

type groupCache struct {
    *redis.Client
    prefix      string
    Initialized bool
}

var Group *groupCache

func newGroup() *groupCache {
    return &groupCache{
        Client: client,
        prefix: "GROUP_",
    }
}

func (c *groupCache) ctx() context.Context {
    return context.Background()
}

func (c *groupCache) key(id string) string {
    return c.prefix + id
}

func (c *groupCache) Initialize(id string, items []string) {
    // remove all members first
    c.Clear(id)
    // add group member cache
    members := make([]interface{}, len(items))
    for i, item := range items {
        members[i] = item
    }
    c.SAdd(c.ctx(), c.key(id), members...)
}

// Exists group cache exists
func (c *groupCache) Exists(id string) bool {
    return c.Client.Exists(c.ctx(), c.key(id)).Val() > 0
}

// MembersCount get group cached members count
func (c *groupCache) MembersCount(id string) int64 {
    return c.SCard(c.ctx(), c.key(id)).Val()
}

// Clear group cache delete
func (c *groupCache) Clear(id string) {
    c.Del(c.ctx(), c.key(id))
}

// MemberExists if member already in group cache
func (c *groupCache) MemberExists(id string, memberID string) bool {
    return c.SIsMember(c.ctx(), c.key(id), memberID).Val()
}

// Add member to group cache
func (c *groupCache) Add(id string, memberID string) {
    // if member not in cache
    if !c.MemberExists(id, memberID) {
        c.SAdd(c.ctx(), c.key(id), memberID)
    }
}

// Remove member from group cache
func (c *groupCache) Remove(id string, memberID string) {
    if c.MemberExists(id, memberID) {
        c.SRem(c.ctx(), c.key(id), memberID)
    }
}

// Members of group
func (c *groupCache) Members(id string) []string {
    return c.SMembers(c.ctx(), c.key(id)).Val()
}
