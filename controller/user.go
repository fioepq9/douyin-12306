package controller

import (
	"douyin-12306/logger"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	var (
		req UserRegisterRequest
		err error
	)

	err = c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{
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

	info, err := service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{
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

	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: Response{
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

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
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
