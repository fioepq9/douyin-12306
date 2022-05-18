package models

type VideoFavorite struct {
	Id      int64
	VideoId int64
	userId  int64
}

func (VideoFavorite) TableName() string {
	return "video_favorite"
}
