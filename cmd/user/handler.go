package main

import (
	"context"
	"douyin-12306/cmd/user/service"
	"douyin-12306/kitex_gen/userKitex"
	"douyin-12306/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *userKitex.UserRegisterRequest) (*userKitex.UserRegisterResponse, error) {
	loginInfo, err := service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		return &userKitex.UserRegisterResponse{
			Response: &userKitex.Response{
				StatusCode: errno.ConvertErr(err).ErrCode,
				StatusMsg:  errno.ConvertErr(err).ErrMsg,
			},
		}, nil
	}
	return &userKitex.UserRegisterResponse{
		Response: &userKitex.Response{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  errno.Success.ErrMsg,
		},
		LoginInfo: loginInfo,
	}, nil
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *userKitex.UserLoginRequest) (*userKitex.UserLoginResponse, error) {
	loginInfo, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		return &userKitex.UserLoginResponse{
			Response: &userKitex.Response{
				StatusCode: errno.ConvertErr(err).ErrCode,
				StatusMsg:  errno.ConvertErr(err).ErrMsg,
			},
		}, nil
	}
	return &userKitex.UserLoginResponse{
		Response: &userKitex.Response{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  errno.Success.ErrMsg,
		},
		LoginInfo: loginInfo,
	}, nil
}

// QueryUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUserInfo(ctx context.Context, req *userKitex.UserInfoRequest) (*userKitex.UserInfoResponse, error) {
	user, err := service.NewUserInfoService(ctx).QueryUserInfo(req)
	if err != nil {
		return &userKitex.UserInfoResponse{
			Response: &userKitex.Response{
				StatusCode: errno.ConvertErr(err).ErrCode,
				StatusMsg:  errno.ConvertErr(err).ErrMsg,
			},
		}, nil
	}
	return &userKitex.UserInfoResponse{
		Response: &userKitex.Response{
			StatusCode: errno.Success.ErrCode,
			StatusMsg:  errno.Success.ErrMsg,
		},
		User: user,
	}, nil
}
