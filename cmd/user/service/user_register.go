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

type UserRegisterService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{ctx: ctx}
}

func (s *UserRegisterService) RegisterUser(req *userKitex.UserRegisterRequest) (loginInfo *userKitex.LoginInfo, err error) {
	rdsCli := rds.NewRedisInstance(dal.D.Redis, s.ctx)
	dbCli := db.NewDB(dal.D.MySQL, s.ctx)

	us := NewUserServiceInstance()

	_, err = rdsCli.QueryUserByUsername(req.Username)
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.L.Errorf("Redis QueryUserByUsername Error: %s", err)
	}
	if err == nil {
		return nil, errno.UserAlreayExistErr
	}
	dbuser, err := dbCli.QueryUserByUsername(req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errno.ConvertErr(err)
	}
	if err == nil {
		err = rdsCli.CacheUserByUsername(pack.NewUserFromDB(dbuser).ToRDS())
		if err != nil {
			logger.L.Errorf("Redis CacheUserByUsername Error: %s", err)
		}
		return nil, errno.UserAlreayExistErr
	}
	sonyId, err := us.sonyflake.NextID()
	if err != nil {
		return nil, errno.SonyflakeGenerateIdFailErr
	}
	var defaultName string
	if len(req.Username) > us.nameMaxLength {
		defaultName = req.Username[:us.nameMaxLength]
	} else {
		defaultName = req.Username
	}
	user := pack.User{
		Id:       int64(sonyId),
		Username: req.Username,
		Password: req.Password,
		Name:     defaultName,
	}
	err = rdsCli.CacheUserByUsername(user.ToRDS())
	if err != nil {
		logger.L.Errorf("Redis CacheUserByUsername Error: %s", err)
	}
	err = dbCli.CreateUser(user.ToDB())
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	token := us.newToken()
	err = rdsCli.CacheUserIdByToken(user.Id, token)
	if err != nil {
		logger.L.Errorf("Redis CacheUserIdByToken Error: %s", err)
	}
	return &userKitex.LoginInfo{
		UserId: user.Id,
		Token:  token,
	}, nil
}
