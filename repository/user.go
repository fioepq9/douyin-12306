package repository

import (
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"sync"
)

type User struct {
	Id       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
}

func (User) TableName() string {
	return "user"
}

type UserDAO struct {
	snowflake *sonyflake.Sonyflake
}

var (
	userDAO  *UserDAO
	userOnce sync.Once
)

func NewUserDAOInstance() *UserDAO {
	userOnce.Do(func() {
		userDAO = &UserDAO{
			snowflake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		}
	})
	return userDAO
}

func (d *UserDAO) Register(username string, password string, name string) (insertUser *User, err error) {
	var user User
	tx := R.MySQL.Table(User{}.TableName()).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// check user exist
	err = tx.Where(&User{Username: username}).Find(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if user.Id != 0 {
		tx.Rollback()
		return nil, errors.New("Found a user with the same username")
	}

	// generate ID
	snowflakeId, err := d.snowflake.NextID()
	if err != nil {
		tx.Rollback()
		return nil, errors.New("snowflake generate ID fail")
	}
	id := int64(snowflakeId)

	// Insert user
	insertUser = &User{
		Id:       id,
		Username: username,
		Password: password,
		Name:     name,
	}
	err = tx.Create(insertUser).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return insertUser, tx.Commit().Error
}