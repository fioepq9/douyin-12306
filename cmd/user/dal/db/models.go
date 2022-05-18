package db

type User struct {
	Id            int64  `gorm:"column:id"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (User) TableName() string {
	return "user"
}

type UserFollow struct {
	Id         int64 `gorm:"column:id"`
	UserId     int64 `gorm:"column:user_id"`
	FollowerId int64 `gorm:"column:follower_id"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}

type UserRelation struct {
	User
	IsFollow bool `gorm:"column:is_follow"`
}
