package service

import (
	"douyin-12306/repository"
	"github.com/pkg/errors"
)

const (
	usernameMaxLength = 32
	passwordMaxLength = 32
	nameMaxLength     = 20
)

type RegisterInfo struct {
	Id    int64
	Token string
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

func NewRegisterInfoFlow(username string, password string) *RegisterInfoFlow {
	return &RegisterInfoFlow{
		username: username,
		password: password,
	}
}

type RegisterInfoFlow struct {
	username string
	password string

	registerInfo *RegisterInfo

	userId int64
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

	id, err := repository.NewUserDAOInstance().Register(f.username, f.password, defaultName)
	if err != nil {
		return err
	}

	f.userId = id

	return nil
}

// packInfo 将 repository 层返回的数据包装，及进行其他逻辑操作
func (f *RegisterInfoFlow) packInfo() error {
	token := f.username + f.password

	f.registerInfo = &RegisterInfo{
		Id:    f.userId,
		Token: token,
	}
	return nil
}
