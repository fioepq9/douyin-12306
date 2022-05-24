package service

import (
	"douyin-12306/pkg/util"
	"douyin-12306/repository"
	"douyin-12306/responses"
	"github.com/gin-gonic/gin"
	"sync"
)

type FeedService struct {
}

var (
	feedService *FeedService
	feedOnce    sync.Once
)

func NewFeedServiceInstance() *FeedService {
	feedOnce.Do(func() {
		feedService = &FeedService{}
	})
	return feedService
}

func (feedService *FeedService) GetFeed(ctx *gin.Context, latestTime int64) (*responses.VideoInfo, error) {

	var (
		id       int64
		nextTime int64
	)

	user := util.GetUser(ctx)
	if user == nil {
		id = -1
	} else {
		id = util.GetUser(ctx).Id
	}

	videoList, err := repository.NewVideoDAOInstance().GetFeed(ctx, latestTime, id)
	if err != nil {
		return nil, err
	}
	if len(videoList) != 0 {
		nextTime = videoList[len(videoList)-1].PublishTime
	} else {
		nextTime = latestTime
	}
	videoInfo := responses.VideoInfo{
		VideoList: videoList,
		NextTime:  nextTime,
	}
	return &videoInfo, err
}
