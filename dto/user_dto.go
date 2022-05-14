package dto

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// UserSimpleDTO 相当于User去掉Password，FollowCount，FollowerCount，用于存储到redis中
type UserSimpleDTO struct {
	Id       int64  `gorm:"column:id" json:"id,omitempty"`
	Username string `gorm:"column:username" json:"username,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
}

func (u *UserSimpleDTO) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *UserSimpleDTO) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
