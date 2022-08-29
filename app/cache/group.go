package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
)

type groupCache struct {
    *redis.Client
    prefix string
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

func (c *groupCache) Init(id string, items []string) {
    // remove all members first
    c.Clear(id)
    // add group member cache
    members := make([]interface{}, len(items))
    for i, item := range items {
        members[i] = item
    }
    c.SAdd(c.ctx(), c.key(id), members...)
}

// Clear group cache delete
func (c *groupCache) Clear(id string) {
    c.Del(c.ctx(), c.key(id))
}

// Exists if member already in group cache
func (c *groupCache) Exists(id string, memberID string) bool {
    return c.SIsMember(c.ctx(), c.key(id), memberID).Val()
}

// Add add member to group cache
func (c *groupCache) Add(id string, memberID string) {
    if !c.Exists(id, memberID) {
        c.SAdd(c.ctx(), c.key(id), memberID)
    }
}

// Remove remove member from group cache
func (c *groupCache) Remove(id string, memberID string) {
    if c.Exists(id, memberID) {
        c.SRem(c.ctx(), c.key(id), memberID)
    }
}
