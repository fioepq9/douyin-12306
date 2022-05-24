package responses

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SuccessResponse(msg string) Response {
	return Response{
		StatusCode: 0,
		StatusMsg:  msg,
	}
}

func ErrorResponse(err error) Response {
	return Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}

type User struct {
	Id            int64  `gorm:"column:user_id" json:"id"`
	Name          string `gorm:"column:username" json:"name"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow"`
}

type LoginInfo struct {
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

type Video struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	Author        User   `gorm:"embedded" json:"author"`
	PlayUrl       string `gorm:"column:play_url" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"column:is_favorite" json:"is_favorite,omitempty"`
	PublishTime   int64  `gorm:"column:publish_time" json:"-"`
}

type VideoInfo struct {
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}
