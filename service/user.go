package service

import (
	"context"
	"douyin-12306/models"
	"douyin-12306/repository"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

const (
	usernameMaxLength = 32
	passwordMaxLength = 32
	nameMaxLength     = 20
	tokenPrefix       = "token:"
)

type RegisterInfo struct {
	Id    int64
	Token string
}

type RegisterInfoFlow struct {
	ctx      context.Context
	username string
	password string

	info *RegisterInfo

	user *models.User
}

type UserDTO struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func NewUserDTO() *UserDTO {
	return &UserDTO{
		Id:       0,
		Username: "",
		Name:     "",
	}
}

func NewRegisterInfoFlow(ctx context.Context, username string, password string) *RegisterInfoFlow {
	return &RegisterInfoFlow{
		ctx:      ctx,
		username: username,
		password: password,
	}
}

func Register(ctx context.Context, username string, password string) (*RegisterInfo, error) {
	flow := NewRegisterInfoFlow(ctx, username, password)
	flowProcessor := FlowProcessor{flow}
	err := flowProcessor.Do()
	if err != nil {
		return nil, err
	}
	return flow.info, nil
}

// 创建token
func newToken() string {
	return tokenPrefix + uuid.NewV4().String()
}

// checkParam
func (f *RegisterInfoFlow) checkParam() error {
	if len(f.username) > usernameMaxLength {
		return errors.New("the username is too long")
	}
	if len(f.password) > passwordMaxLength {
		return errors.New("the password is too long")
	}
	return nil
}

// prepareInfo 构建参数, 向 repository 层请求数据
func (f *RegisterInfoFlow) prepareInfo() (err error) {
	var defaultName string
	if len(f.username) > nameMaxLength {
		defaultName = f.username[:nameMaxLength]
	} else {
		defaultName = f.username
	}

	user, err := repository.NewUserDAOInstance().Register(f.ctx, f.username, f.password, defaultName)
	if err != nil {
		return err
	}

	f.user = user

	return nil
}

// packInfo 将 repository 层返回的数据包装，及进行其他逻辑操作
func (f *RegisterInfoFlow) packInfo() error {
	// 1.获取redis连接客户端
	// redisClient := redisUtil.NewRedisClient()

	// 2.生成用户token
	token := newToken()

	// 3.将用户信息（转为DTO）存储到redis中
	// user := f.user
	// 3.1 User转UserDTO
	// userDTO := NewUserDTO()
	// err := copier.Copy(userDTO, user)
	// if err != nil {
	// 	return err
	// }

	// 3.2 存入redis
	// err = redisUtil.RedisSetSimpleStruct(redisClient, token, userDTO)
	err := repository.R.Redis.Set(f.ctx, token, f.user, f.user.Expiration()).Err()
	if err != nil {
		return err
	}

	// 4. 将返回信息放置到registerInfo
	f.info = &RegisterInfo{
		Id:    f.user.Id,
		Token: token,
	}
	return nil
}
