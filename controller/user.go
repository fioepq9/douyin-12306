package controller

import (
	"douyin-12306/logger"
	"douyin-12306/requests"
	"douyin-12306/responses"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var (
		req requests.UserRegisterRequest
		err error
	)

	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, responses.UserRegisterResponse{
			Response: responses.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	logger.L.Debug("Register 接口的 Request", map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	})

	info, err := service.NewUserServiceInstance().Register(c, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, responses.UserRegisterResponse{
			Response: responses.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	logger.L.Debug("Register 接口的 Response", map[string]interface{}{
		"user_id": info.Id,
		"token":   info.Token,
	})

	c.JSON(http.StatusOK, responses.UserRegisterResponse{
		Response: responses.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: info.Id,
		Token:  info.Token,
	})
}

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Login(c *gin.Context) {
	var (
		req requests.UserLoginRequest
		err error
	)
	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, responses.UserRegisterResponse{
			Response: responses.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	info, err := service.NewUserServiceInstance().Login(c, req.Username, req.Password)
	// 错误信息
	if err != nil {
		c.JSON(http.StatusOK, responses.UserRegisterResponse{
			Response: responses.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 正确返回
	c.JSON(http.StatusOK, responses.UserRegisterResponse{
		Response: responses.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: info.Id,
		Token:  info.Token,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
