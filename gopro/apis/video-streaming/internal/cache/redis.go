package cache

import (
    "context"
   "github.com/redis/go-redis/v9"
   "time"
)

var RedisClient *redis.Client

func InitRedis(addr, password string) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       0,
    })
}

func CacheVideo(ctx context.Context, key string, value []byte) error {
    return RedisClient.Set(ctx, key, value, 10*time.Second).Err()
}

