package service

import (
	"douyin-12306/repository"
	"douyin-12306/utils/redisUtil"
	"github.com/jinzhu/copier"
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
	username string
	password string

	registerInfo *RegisterInfo

	insertUser *repository.User
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

func NewRegisterInfoFlow(username string, password string) *RegisterInfoFlow {
	return &RegisterInfoFlow{
		username: username,
		password: password,
	}
}

func Register(username string, password string) (*RegisterInfo, error) {
	flow := NewRegisterInfoFlow(username, password)
	flowProcessor := FlowProcessor{flow}
	err := flowProcessor.Do()
	if err != nil {
		return nil, err
	}
	return flow.registerInfo, nil
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

	user, err := repository.NewUserDAOInstance().Register(f.username, f.password, defaultName)
	if err != nil {
		return err
	}

	f.insertUser = user

	return nil
}

// packInfo 将 repository 层返回的数据包装，及进行其他逻辑操作
func (f *RegisterInfoFlow) packInfo() error {
	// 1.获取redis连接客户端
	redisClient := redisUtil.NewRedisClient()

	// 2.生成用户token
	token := newToken()

	// 3.将用户信息（转为DTO）存储到redis中
	user := f.insertUser
	// 3.1 User转UserDTO
	userDTO := NewUserDTO()
	err := copier.Copy(userDTO, user)
	if err != nil {
		return err
	}

	// 3.2 存入redis
	err = redisUtil.RedisSetSimpleStruct(redisClient, token, userDTO)
	if err != nil {
		return err
	}

	// 4. 将返回信息放置到registerInfo
	f.registerInfo = &RegisterInfo{
		Id:    f.insertUser.Id,
		Token: token,
	}
	return nil
}
