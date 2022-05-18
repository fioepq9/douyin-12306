package handlers

import (
	"douyin-12306/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// SendErrorResponse pack error response
func SendErrorResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

type LoginInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
