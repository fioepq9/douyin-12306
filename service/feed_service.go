package service

import (
	"douyin-12306/responses"
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

func (feedService *FeedService) GetFeed(latestTime int64, token string) ([]responses.Video, error) {

	return nil, nil
}
