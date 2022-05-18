package controller

import (
	"douyin-12306/logger"
	"douyin-12306/responses"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []responses.Video `json:"video_list,omitempty"`
	NextTime  int64             `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var (
		latestTime int64
		err        error
	)

	latestTimeStr := c.Query("latest_time")
	token := c.Query("token")

	logger.L.Debug("feed 接口的 Request", map[string]interface{}{
		"latestTime": latestTimeStr,
		"token":      token,
	})

	if latestTimeStr == "" {
		latestTime = time.Now().Unix()
	} else {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 1, StatusMsg: "时间参数格式错误"},
			})
			return
		}
	}

	//var videoList
	videoList, err := service.NewFeedServiceInstance().GetFeed(latestTime, token)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
