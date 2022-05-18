package pack

import (
	"douyin-12306/cmd/user/dal/db"
	"douyin-12306/cmd/user/dal/rds"
	"douyin-12306/kitex_gen/userKitex"
)

type User struct {
	Id            int64
	Username      string
	Password      string
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

func NewUserFromDB(u *db.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:            u.Id,
		Username:      u.Username,
		Password:      u.Password,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}
}

func (u *User) ToDB() *db.User {
	if u == nil {
		return nil
	}
	return &db.User{
		Id:            u.Id,
		Username:      u.Username,
		Password:      u.Password,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}
}

func NewUserFromDBR(u *db.UserRelation) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:            u.Id,
		Username:      u.Username,
		Password:      u.Password,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func (u *User) ToDBR() *db.UserRelation {
	if u == nil {
		return nil
	}
	return &db.UserRelation{
		User: db.User{
			Id:            u.Id,
			Username:      u.Username,
			Password:      u.Password,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
		},
		IsFollow: u.IsFollow,
	}
}

func NewUserFromRDS(u *rds.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:            u.Id,
		Username:      u.Username,
		Password:      u.Password,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}
}

func (u *User) ToRDS() *rds.User {
	if u == nil {
		return nil
	}
	return &rds.User{
		Id:            u.Id,
		Username:      u.Username,
		Password:      u.Password,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}
}

func (u *User) ToResp() *userKitex.User {
	if u == nil {
		return nil
	}
	return &userKitex.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
