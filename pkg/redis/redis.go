package redis

import (
	"admin-go-api/common/config"
	"context"
	"github.com/go-redis/redis/v8"
)

// redis 初始化
var (
	RedisDb *redis.Client
)

// 初始化连接
func SetupRedisDB() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
