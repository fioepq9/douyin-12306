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
			Response: responses.Response{StatusCode: 1, StatusMsg: "请求参数格式错误"},
		})
		return
	}

	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}

	//var videoList
	videoInfo, err := service.NewFeedServiceInstance().GetFeed(c, req.LatestTime)
	if err != nil {
		c.JSON(
			http.StatusOK,
			responses.FeedResponse{
				Response: responses.Response{StatusCode: 1, StatusMsg: "获取视频资源失败"},
			},
		)
		return
	}
	var message string
	if len(videoInfo.VideoList) == 0 {
		message = "暂无更多视频资源！"
	} else {
		message = "获取视频资源成功！"
	}
	c.JSON(
		http.StatusOK,
		responses.FeedResponse{
			Response:  responses.Response{StatusCode: 0, StatusMsg: message},
			VideoInfo: *videoInfo,
		},
	)
}
