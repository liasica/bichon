package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
    "time"
)

var client *redis.Client

var (
    Set func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
    Get func(ctx context.Context, key string) *redis.StringCmd
    Del func(ctx context.Context, keys ...string) *redis.IntCmd

    HSet    func(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
    HDel    func(ctx context.Context, key string, fields ...string) *redis.IntCmd
    HGet    func(ctx context.Context, key, field string) *redis.StringCmd
    HGetAll func(ctx context.Context, key string) *redis.StringStringMapCmd

    Expire func(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd

    LLen   func(ctx context.Context, key string) *redis.IntCmd
    RPush  func(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
    LIndex func(ctx context.Context, key string, index int64) *redis.StringCmd
    LRange func(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
)

func CreateClient(addr, password string, db int) {
    client = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    Set = client.Set
    Get = client.Get
    Del = client.Del

    HSet = client.HSet
    HDel = client.HDel
    HGet = client.HGet
    HGetAll = client.HGetAll

    Expire = client.Expire

    LLen = client.LLen
    RPush = client.RPush
    LIndex = client.LIndex
    LRange = client.LRange

    Group = newGroup()
}
