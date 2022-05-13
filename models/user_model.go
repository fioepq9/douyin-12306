package models

import (
	"bytes"
	"douyin-12306/pkg/util"
	jsoniter "github.com/json-iterator/go"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	Username      string `gorm:"column:username" json:"username,omitempty"`
	Password      string `gorm:"column:password" json:"password,omitempty"`
	Name          string `gorm:"column:name" json:"name,omitempty"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (User) Expiration() time.Duration {
	return time.Hour
}

func (User) KeyPrefix() string {
	return "user:"
}

func (User) UsernameKeyPrefix() string {
	return "username:"
}

func (u *User) Key() string {
	buf := bytes.NewBufferString(u.KeyPrefix())
	buf.WriteString(util.Int64ToString(u.Id))
	return buf.String()
}

func (u *User) UsernameKey() string {
	buf := bytes.NewBufferString(u.UsernameKeyPrefix())
	buf.WriteString(u.Username)
	return buf.String()
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
