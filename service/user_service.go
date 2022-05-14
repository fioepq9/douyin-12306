package service

import (
	"context"
	"douyin-12306/dto"
	"douyin-12306/pkg/util"
	"douyin-12306/repository"
	"douyin-12306/responses"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"sync"
)

const (
	nameMaxLength = 20
	TokenPrefix   = "token:"
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

func (s *UserService) Register(ctx context.Context, username, password string) (info *responses.LoginInfo, err error) {
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

	info = &responses.LoginInfo{
		Id:    user.Id,
		Token: s.newToken(),
	}

	// User转UserSimpleDTO
	userDTO := &dto.UserSimpleDTO{}
	err = copier.Copy(userDTO, user)
	if err != nil {
		return nil, err
	}
	// 存储userDTO
	err = repository.R.Redis.Set(ctx, TokenPrefix+info.Token, userDTO, user.Expiration()).Err()
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Login 登录
func (s *UserService) Login(ctx context.Context, username string, password string) (info *responses.LoginInfo, err error) {
	// 1.查询用户是否存在
	user, err := repository.NewUserDAOInstance().GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	// 2.校验密码
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	// 3.登录成功
	// 3.1 User转为UserDTO
	userDTO := &dto.UserSimpleDTO{}
	err = copier.Copy(userDTO, user)
	if err != nil {
		return nil, err
	}

	// 3.2 存到redis
	token := s.newToken()
	err = repository.R.Redis.Set(ctx, TokenPrefix+token, userDTO, user.Expiration()).Err()
	if err != nil {
		return nil, err
	}

	// 3.3 返回
	return &responses.LoginInfo{
		Id:    user.Id,
		Token: token,
	}, nil
}

func (s *UserService) GetUserInfo(c *gin.Context, selectId int64) (*responses.User, error) {
	// 1.查询selectId对应用户信息
	user := repository.NewUserDAOInstance().GetUserByUserId(c, selectId)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	// 用户存在时将User转为UserDTO
	var userDTO = &dto.UserDTO{}
	err := copier.Copy(userDTO, user)
	if err != nil {
		return nil, errors.New("User到UserDTO转化失败")
	}

	// 2.查询isFollow信息
	userId := util.GetUser(c).Id
	isFollow := repository.NewUserDAOInstance().IsUserFollow(c, userId, selectId)

	// 3.组装到结果返回
	var userResponse = responses.User{}
	err = copier.Copy(userResponse, userDTO)
	if err != nil {
		return nil, errors.New("UserDTO到responses.User转化失败")
	}
	userResponse.IsFollow = isFollow

	return &userResponse, nil
}
