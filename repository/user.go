package repository

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"sync"
)

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

type UserDAO struct {
	sonyflake *sonyflake.Sonyflake
}

var (
	userDAO  *UserDAO
	userOnce sync.Once
)

func NewUserDAOInstance() *UserDAO {
	userOnce.Do(func() {
		userDAO = &UserDAO{
			sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		}
	})
	return userDAO
}

func (d *UserDAO) Register(username string, password string, name string) (*User, error) {
	var user User
	tx := R.MySQL.Table(User{}.TableName()).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check：该username是否已存在
	err := tx.Where(&User{Username: username}).Take(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			err = errors.New("用户名已存在")
		}
		tx.Rollback()
		return nil, err
	}

	// 使用 sonyflake 生成 ID
	sonyId, err := d.sonyflake.NextID()
	if err != nil {
		tx.Rollback()
		return nil, errors.New("sonyflake generate ID fail")
	}
	id := int64(sonyId)

	// Insert user
	user = User{
		Id:       id,
		Username: username,
		Password: password,
		Name:     name,
	}
	err = tx.Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, tx.Commit().Error
}
