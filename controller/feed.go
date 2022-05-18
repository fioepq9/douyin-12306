package controller

import (
	"douyin-12306/requests"
	"douyin-12306/responses"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	req := requests.FeedRequest{}
	err := c.BindQuery(&req)

	if err != nil {
		c.JSON(http.StatusOK, responses.FeedResponse{
			Response: responses.Response{StatusCode: 1, StatusMsg: "参数格式错误"},
		})
	}

	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}

	//var videoList
	videoInfo, err := service.NewFeedServiceInstance().GetFeed(c, req.LatestTime)
	c.JSON(
		http.StatusOK,
		responses.FeedResponse{
			Response:  responses.Response{StatusCode: 0},
			VideoInfo: *videoInfo,
		},
	)
}
