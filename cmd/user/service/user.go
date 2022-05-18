package service

import (
	"sync"

	"github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
)

type UserService struct {
	sonyflake     *sonyflake.Sonyflake
	nameMaxLength int
	newToken      func() string
}

var (
	userService *UserService
	userOnce    sync.Once
)

func NewUserServiceInstance() *UserService {
	userOnce.Do(func() {
		userService = &UserService{
			sonyflake:     sonyflake.NewSonyflake(sonyflake.Settings{}),
			nameMaxLength: 20,
			newToken: func() string {
				return uuid.NewV4().String()
			},
		}
	})
	return userService
}
