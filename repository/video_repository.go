package repository

import (
	"context"
	"douyin-12306/dto"
	"douyin-12306/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
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

func (videoDAO *VideoDAO) GetFeed(ctx context.Context, latestTime int64, userId int64) ([]dto.VideoDTO, error) {
	var (
		videoDtoList []dto.VideoDTO
		userIdStr    string
		err          error
	)
	db := R.MySQL
	db.Preload("User").Find(&dto.VideoDTO{})

	userIdStr = strconv.FormatInt(userId, 10)

	columns := [...]string{"id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "publish_time", "user.id user_id", "`name` username",
		"follow_count", "follower_count", "IF(video_favorite.id,1,0) is_favorite", "IF(user_follow.id,1,0) is_follow"}
	err = db.WithContext(ctx).Table(models.Video{}.TableName()).Select(columns).
		Joins("LEFT JOIN user ON author_id=user.id").
		Joins("LEFT JOIN video_favorite ON video_favorite.video_id=video.id AND video_favorite.user_id="+userIdStr).
		Joins("LEFT JOIN user_follow ON user_follow.user_id=author_id AND follower_id="+userIdStr).
		Where("publish_time<?", latestTime).
		Order("publis_time DESC").
		Limit(30).
		Take(&videoDtoList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("视频不存在")
		}
		return nil, err
	}
	return videoDtoList, nil
}
