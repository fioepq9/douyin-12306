package repo

import (
	"context"
	"douyin-12306/config"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (rds *redis.Client, err error) {
	redisURL := fmt.Sprintf("redis://%s:%s@%s/%s",
		config.C.Redis.Username,
		config.C.Redis.Password,
		config.C.Redis.Addr,
		config.C.Redis.DB,
	)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	rds = redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = rds.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return rds, nil
}
