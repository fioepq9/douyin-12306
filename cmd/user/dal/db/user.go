package db

import (
	"context"

	"gorm.io/gorm"
)

type MySQL struct {
	db  *gorm.DB
	ctx context.Context
}

func NewDB(db *gorm.DB, ctx context.Context) *MySQL {
	return &MySQL{
		db:  db,
		ctx: ctx,
	}
}

func (s *MySQL) CreateUser(user *User) error {
	return s.db.WithContext(s.ctx).Table(user.TableName()).Create(user).Error
}

func (s *MySQL) QueryUserByUserId(userId int64) (*User, error) {
	user := User{
		Id: userId,
	}
	err := s.db.WithContext(s.ctx).Table(
		user.TableName()).Where(
		user).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MySQL) QueryUserByUsername(username string) (*User, error) {
	user := User{
		Username: username,
	}
	err := s.db.WithContext(s.ctx).Table(
		user.TableName()).Where(
		user).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MySQL) QueryUserRelationByUserId(selectUserId, curUserId int64) (*UserRelation, error) {
	var user UserRelation
	err := s.db.WithContext(s.ctx).Table(User{}.TableName()).Select("*, IF(EXISTS(?), TRUE, FALSE) as is_follow",
		s.db.Table(UserFollow{}.TableName()).Where(UserFollow{
			UserId:     selectUserId,
			FollowerId: curUserId,
		})).Where(User{
		Id: selectUserId,
	}).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
