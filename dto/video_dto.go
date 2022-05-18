package dto

import "douyin-12306/responses"

// 用来接从数据库中查出的字段,比response中的video要多一个publisTime
type VideoDTO struct {
	Id            int64          `gorm:"column:id"`
	Author        responses.User `gorm:"foreignkey:Id;references:user_id"`
	PlayUrl       string         `gorm:"column:play_url"`
	CoverUrl      string         `gorm:"column:cover_url,omitempty"`
	FavoriteCount int64          `gorm:"column:favorite_count,omitempty"`
	CommentCount  int64          `gorm:"column:comment_count,omitempty"`
	IsFavorite    bool           `gorm:"column:is_favorite,omitempty"`
	PublishTime   int64          `gorm:"column:publish_time"`
}
