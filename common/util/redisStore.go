// redis存取验证码配置
package util

import (
	"admin-go-api/common/constant"
	"admin-go-api/pkg/redis"
	"context"
	"log"
	"time"
)

var ctx = context.Background()

type RedisStore struct {
}

// 存验证码
func (r *RedisStore) Set(id string, value string) error {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
	}
	return err
}

// 取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	value, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		log.Panicln(err.Error())
	}
	return value
}

// 验证验证码
func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	key := constant.LOGIN_CODE + id
	code, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		log.Panicln(err.Error())
	}
	if code == answer {
		return true
	}
	return false

}
