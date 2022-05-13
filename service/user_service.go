package service

import (
	"context"
	"douyin-12306/repository"
	"github.com/satori/go.uuid"
	"sync"
)

const (
	nameMaxLength = 20
	tokenPrefix   = "token:"
)

type UserService struct {
	newToken func() string
}

var (
	userService *UserService
	userOnce    sync.Once
)

func NewUserServiceInstance() *UserService {
	userOnce.Do(func() {
		userService = &UserService{
			newToken: func() string {
				return uuid.NewV4().String()
			},
		}
	})
	return userService
}

type RegisterInfo struct {
	Id    int64
	Token string
}

func (s *UserService) Register(ctx context.Context, username, password string) (info *RegisterInfo, err error) {
	var defaultName string
	if len(username) > nameMaxLength {
		defaultName = username[:nameMaxLength]
	} else {
		defaultName = username
	}

	user, err := repository.NewUserDAOInstance().Register(ctx, username, password, defaultName)
	if err != nil {
		return nil, err
	}

	info = &RegisterInfo{
		Id:    user.Id,
		Token: s.newToken(),
	}

	err = repository.R.Redis.Set(ctx, tokenPrefix+info.Token, user, user.Expiration()).Err()
	if err != nil {
		return nil, err
	}

	return info, nil
}
