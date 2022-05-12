package redisUtil

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "120.25.0.102:6379",
		Password: "root",
		DB:       0,
	})
}

func RedisSetSimpleStruct(client *redis.Client, key string, structData interface{}) error {
	ctx := context.Background()
	marshal, err := json.Marshal(structData)
	if err != nil {
		return err
	}
	client.Set(ctx, key, string(marshal), 0)
	return nil
}

func RedisGetSimpleStruct(client *redis.Client, key string, out interface{}) error {
	ctx := context.Background()
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), out)
	if err != nil {
		return err
	}

	return nil
}
