package service

import (
	"context"
	"douyin-12306/cmd/user/dal"
	"douyin-12306/cmd/user/dal/db"
	"douyin-12306/cmd/user/dal/rds"
	"douyin-12306/cmd/user/pack"
	"douyin-12306/kitex_gen/userKitex"
	"douyin-12306/pkg/errno"
	"errors"

	"github.com/go-redis/redis/v8"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) QueryUserInfo(req *userKitex.UserInfoRequest) (*userKitex.User, error) {
	rdsCli := rds.NewRedisInstance(dal.D.Redis, s.ctx)
	dbCli := db.NewDB(dal.D.MySQL, s.ctx)

	id, err := rdsCli.QueryUserIdByToken(req.Token)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, errno.ConvertErr(err)
	}
	if errors.Is(err, redis.Nil) {
		return nil, errno.UserNotLoginErr
	}
	userDBR, err := dbCli.QueryUserRelationByUserId(req.UserId, id)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}
	return pack.NewUserFromDBR(userDBR).ToResp(), nil
}
