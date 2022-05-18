package service

import (
	"context"
	"douyin-12306/cmd/user/dal"
	"douyin-12306/cmd/user/dal/db"
	"douyin-12306/cmd/user/dal/rds"
	"douyin-12306/cmd/user/pack"
	"douyin-12306/kitex_gen/userKitex"
	"douyin-12306/logger"
	"douyin-12306/pkg/errno"
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserLoginService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *UserLoginService {
	return &UserLoginService{ctx: ctx}
}

func (s *UserLoginService) LoginUser(req *userKitex.UserLoginRequest) (loginInfo *userKitex.LoginInfo, err error) {
	rdsCli := rds.NewRedisInstance(dal.D.Redis, s.ctx)
	dbCli := db.NewDB(dal.D.MySQL, s.ctx)

	var user *pack.User

	userRDS, err := rdsCli.QueryUserByUsername(req.Username)
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.L.Errorf("Redis QueryUserByUsername Error: %s", err)
	} else if errors.Is(err, redis.Nil) {
		var userDB *db.User
		userDB, err = dbCli.QueryUserByUsername(req.Username)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ConvertErr(err)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotExistErr
		}
		user = pack.NewUserFromDB(userDB)
		err = rdsCli.CacheUserByUsername(user.ToRDS())
		if err != nil {
			logger.L.Errorf("Redis CacheUserByUsername Error: %s", err)
		}
	} else {
		user = pack.NewUserFromRDS(userRDS)
	}
	if user.Password != req.Password {
		return nil, errno.LoginErr
	}
	token := NewUserServiceInstance().newToken()
	err = rdsCli.CacheUserIdByToken(user.Id, token)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	return &userKitex.LoginInfo{
		UserId: user.Id,
		Token:  token,
	}, nil
}
