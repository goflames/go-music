package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

// 初始化 Redis 客户端
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis 服务器地址
	DB:   0,                // Redis 数据库号
})

// SaveTokenToBlacklist 将 JWT 添加到 Redis 黑名单，设置过期时间
func SaveTokenToBlacklist(token string, exp time.Duration) error {
	return rdb.Set(ctx, token, "blacklisted", exp).Err()
}

// IsTokenInBlacklist 检查 JWT 是否在 Redis 黑名单中
func IsTokenInBlacklist(token string) (bool, error) {
	result, err := rdb.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil // 不存在于黑名单
	} else if err != nil {
		return false, err
	}
	return result == "blacklisted", nil
}

// SaveTokenToRedis 将 JWT 保存到 Redis，并设置 24 小时的过期时间
func SaveTokenToRedis(token string, userId uint) error {
	expiration := time.Hour * 24 // 24 小时过期时间
	return rdb.Set(ctx, token, userId, expiration).Err()
}
