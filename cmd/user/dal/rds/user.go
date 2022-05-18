package rds

import (
	"context"
	"douyin-12306/pkg/errno"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisInstance(client *redis.Client, ctx context.Context) *Redis {
	return &Redis{
		client: client,
		ctx:    ctx,
	}
}

func (r *Redis) QueryUserByUserId(userId int64) (*User, error) {
	user := User{Id: userId}
	err := r.client.Get(r.ctx, user.Key()).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Redis) QueryUserByUsername(username string) (*User, error) {
	user := User{Username: username}
	err := r.client.Get(r.ctx, user.UsernameKey()).Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return r.QueryUserByUserId(user.Id)
}

func (r *Redis) QueryUserIdByToken(token string) (int64, error) {
	var id int64
	t := Token{token: token}
	err := r.client.Get(r.ctx, t.Key()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Redis) CacheUserByUserId(user *User) error {
	if user == nil {
		return errno.ParamErr
	}
	return r.client.Set(r.ctx, user.Key(), user, user.Expiration()).Err()
}

func (r *Redis) CacheUserByUsername(user *User) error {
	if user == nil {
		return errno.ParamErr
	}
	err := r.client.Set(r.ctx, user.UsernameKey(), user.Id, user.Expiration()).Err()
	if err != nil {
		return err
	}
	return r.CacheUserByUserId(user)
}

func (r *Redis) CacheUserIdByToken(userId int64, token string) error {
	t := Token{token: token}
	return r.client.Set(r.ctx, t.Key(), userId, t.Expiration()).Err()
}
