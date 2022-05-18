package service

import (
	"douyin-12306/repository"
	"douyin-12306/responses"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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

	//@todo
	//id := util.GetUser(ctx).Id
	videoDtoList, err := repository.NewVideoDAOInstance().GetFeed(ctx, latestTime, 0)
	if err != nil {
		return nil, err
	}
	//nextTime := videoDtoList[len(videoDtoList)-1].PublishTime
	var videoList []responses.Video
	copier.Copy(videoList, videoDtoList)
	videoInfo := responses.VideoInfo{
		VideoList: videoList,
		NextTime:  0,
	}
	return &videoInfo, err
}
