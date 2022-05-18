package handlers

import (
	"douyin-12306/cmd/api/rpc"
	"douyin-12306/logger"
	"douyin-12306/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var param UserRegisterParam

	// 检查参数
	err := c.BindQuery(&param)
	if err != nil {
		SendErrorResponse(c, errno.ParamErr.WithMessage(err.Error()))
		return
	}
	logger.L.Info(param)
	// rpc
	resp, err := rpc.RegisterUser(c, param.ToRequest())
	if err != nil {
		SendErrorResponse(c, errno.ConvertErr(err))
		return
	}

	// success
	c.JSON(http.StatusOK, UserRegisterResponse{}.PackResponse(resp))
}

func LoginUser(c *gin.Context) {
	var param UserLoginParam

	// 检查参数
	err := c.BindQuery(&param)
	if err != nil {
		SendErrorResponse(c, errno.ParamErr.WithMessage(err.Error()))
		return
	}

	// rpc
	resp, err := rpc.LoginUser(c, param.ToRequest())
	if err != nil {
		SendErrorResponse(c, errno.ConvertErr(err))
		return
	}

	// success
	c.JSON(http.StatusOK, UserLoginResponse{}.PackResponse(resp))
}

func QueryUserInfo(c *gin.Context) {
	var param UserInfoParam

	// 检查参数
	err := c.BindQuery(&param)
	if err != nil {
		SendErrorResponse(c, errno.ParamErr.WithMessage(err.Error()))
		return
	}

	// rpc
	resp, err := rpc.QueryUserInfo(c, param.ToRequest())
	if err != nil {
		SendErrorResponse(c, errno.ConvertErr(err))
		return
	}

	// success
	c.JSON(http.StatusOK, UserInfoResponse{}.PackResponse(resp))
}
