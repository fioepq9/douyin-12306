package repository

import (
	"context"
	"douyin-12306/models"
	"douyin-12306/responses"
	"sync"
)

type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAOInstance() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = &VideoDAO{}
	})
	return videoDAO
}

func (videoDAO *VideoDAO) GetFeed(ctx context.Context, latestTime int64, userId int64) ([]responses.Video, error) {
	var (
		videoList []responses.Video
		err       error
	)
	db := R.MySQL

	if userId == -1 {
		columns := "video.id,play_url,cover_url,favorite_count,comment_count,publish_time,`user`.id user_id,`name` username,follow_count,follower_count"
		err = db.WithContext(ctx).Table(models.Video{}.TableName()).Select(columns).
			Joins("LEFT JOIN `user` ON author_id=user.id").
			Where("publish_time<?", latestTime).
			Order("publish_time DESC").
			Limit(30).
			Find(&videoList).Error
	} else {
		columns := "video.id,play_url,cover_url,favorite_count,comment_count,publish_time,`user`.id user_id,`name` username,follow_count,follower_count,IF(video_favorite.id,1,0) is_favorite,IF(user_follow.id,1,0) is_follow"
		err = db.WithContext(ctx).Table(models.Video{}.TableName()).Select(columns).
			Joins("LEFT JOIN `user` ON author_id=user.id").
			Joins("LEFT JOIN video_favorite ON video_favorite.video_id=video.id AND video_favorite.user_id=?", userId).
			Joins("LEFT JOIN user_follow ON user_follow.user_id=author_id AND follower_id=?", userId).
			Where("publish_time<?", latestTime).
			Order("publish_time DESC").
			Limit(30).
			Find(&videoList).Error
	}
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
