package handlers

import "douyin-12306/kitex_gen/userKitex"

type UserRegisterParam struct {
	Username string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

func (p *UserRegisterParam) ToRequest() *userKitex.UserRegisterRequest {
	if p == nil {
		return nil
	}
	return &userKitex.UserRegisterRequest{
		Username: p.Username,
		Password: p.Password,
	}
}

type UserLoginParam struct {
	Username string `form:"username" binding:"required,min=1,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}

func (p *UserLoginParam) ToRequest() *userKitex.UserLoginRequest {
	if p == nil {
		return nil
	}
	return &userKitex.UserLoginRequest{
		Username: p.Username,
		Password: p.Password,
	}
}

type UserInfoParam struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func (p *UserInfoParam) ToRequest() *userKitex.UserInfoRequest {
	if p == nil {
		return nil
	}
	return &userKitex.UserInfoRequest{
		UserId: p.UserId,
		Token:  p.Token,
	}
}

type UserRegisterResponse struct {
	Response
	LoginInfo
}

func (UserRegisterResponse) PackResponse(resp *userKitex.UserRegisterResponse) *UserRegisterResponse {
	if resp == nil {
		return nil
	}
	return &UserRegisterResponse{
		Response: Response{
			StatusCode: resp.Response.StatusCode,
			StatusMsg:  resp.Response.StatusMsg,
		},
		LoginInfo: LoginInfo{
			UserId: resp.LoginInfo.UserId,
			Token:  resp.LoginInfo.Token,
		},
	}
}

type UserLoginResponse struct {
	Response
	LoginInfo
}

func (UserLoginResponse) PackResponse(resp *userKitex.UserLoginResponse) *UserLoginResponse {
	if resp == nil {
		return nil
	}
	return &UserLoginResponse{
		Response: Response{
			StatusCode: resp.Response.StatusCode,
			StatusMsg:  resp.Response.StatusMsg,
		},
		LoginInfo: LoginInfo{
			UserId: resp.LoginInfo.UserId,
			Token:  resp.LoginInfo.Token,
		},
	}
}

type UserInfoResponse struct {
	Response
	User `json:"user"`
}

func (UserInfoResponse) PackResponse(resp *userKitex.UserInfoResponse) *UserInfoResponse {
	if resp == nil {
		return nil
	}
	return &UserInfoResponse{
		Response: Response{
			StatusCode: resp.Response.StatusCode,
			StatusMsg:  resp.Response.StatusMsg,
		},
		User: User{
			Id:            resp.User.Id,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
		},
	}
}
