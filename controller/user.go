package controller

import (
	"douyin-12306/logger"
	"douyin-12306/requests"
	"douyin-12306/responses"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册接口
func Register(c *gin.Context) {
	var (
		req requests.UserRegisterRequest
		err error
	)

	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}
	logger.L.Debugw("Register 接口的 Request", map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	})

	info, err := service.NewUserServiceInstance().Register(c, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}
	logger.L.Debugw("Register 接口的 Response", map[string]interface{}{
		"user_id": info.Id,
		"token":   info.Token,
	})

	c.JSON(http.StatusOK, responses.UserRegisterResponse{
		Response: responses.SuccessResponse("register success"),
		UserId:   info.Id,
		Token:    info.Token,
	})
}

// Login 登录接口
func Login(c *gin.Context) {
	var (
		req requests.UserLoginRequest
		err error
	)
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}

	info, err := service.NewUserServiceInstance().Login(c, req.Username, req.Password)
	// 错误信息
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}

	// 正确返回
	c.JSON(http.StatusOK, responses.UserLoginResponse{
		Response: responses.SuccessResponse("login success"),
		UserId:   info.Id,
		Token:    info.Token,
	})
}

// UserInfo 用户信息接口
func UserInfo(c *gin.Context) {
	var (
		req requests.UserInfoRequest
		err error
	)
	// 绑定参数
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}

	info, err := service.NewUserServiceInstance().GetUserInfo(c, req.UserId)
	if err != nil {
		c.JSON(http.StatusOK, responses.ErrorResponse(err))
		return
	}

	// 正确返回
	c.JSON(http.StatusOK, responses.UserInfoResponse{
		Response: responses.SuccessResponse("query user info success"),
		User: responses.User{
			Id:            info.Id,
			Name:          info.Name,
			FollowCount:   info.FollowCount,
			FollowerCount: info.FollowerCount,
			IsFollow:      info.IsFollow,
		},
	})
}
