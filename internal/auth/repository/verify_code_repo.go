package repository

import (
	"chat/internal/auth/model"
	myredis "chat/internal/auth/repository/go-redis"
	"context"
	"fmt"
	"time"
)

type VerifyCodeRepository struct {
}

// 设置键和设置时间非原子，可能会出现问题，后续需要修改
func (*VerifyCodeRepository) Save(ctx context.Context, key string, ttl, value int) (bool, error) {
	cmd := myredis.Rdb.HSet(ctx, key, model.EmailRedisValue{
		Code:    value,
		Attemps: 0,
	})
	count, err := cmd.Result()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, fmt.Errorf("key already exists")
	}
	myredis.Rdb.Expire(ctx, key, time.Second*time.Duration(ttl))
	return true, nil
}

func (*VerifyCodeRepository) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := myredis.Rdb.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if ttl == -1 {
		return 0, fmt.Errorf("not set ttl")
	}
	if ttl == -2 {
		return 0, fmt.Errorf("key not exists")
	}
	return ttl, nil
}
