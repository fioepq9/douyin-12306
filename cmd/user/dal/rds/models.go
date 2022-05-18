package rds

import (
	"bytes"
	"douyin-12306/pkg/util"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	Id            int64  `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

func (User) Expiration() time.Duration {
	return time.Hour
}

func (u *User) Key() string {
	buf := bytes.NewBufferString("user:")
	buf.WriteString(util.Int64ToString(u.Id))
	return buf.String()
}

func (u *User) UsernameKey() string {
	buf := bytes.NewBufferString("username:")
	buf.WriteString(u.Username)
	return buf.String()
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type Token struct {
	token string
}

func (Token) Expiration() time.Duration {
	return 12 * time.Hour
}

func (t *Token) Key() string {
	buf := bytes.NewBufferString("token:")
	buf.WriteString(t.token)
	return buf.String()
}
