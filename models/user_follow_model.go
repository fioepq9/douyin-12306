package models

type UserFollow struct {
	Id         int64 `gorm:"column:id" json:"id,omitempty"`
	UserId     int64 `gorm:"column:user_id" json:"user_id,omitempty"`
	FollowerId int64 `gorm:"column:follower_id" json:"follower_id,omitempty"`
}

func (UserFollow) TableName() string {
	return "user_follow"
}

func (u *UserFollow) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *UserFollow) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
