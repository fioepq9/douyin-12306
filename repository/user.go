package repository

import (
	"context"
	"douyin-12306/logger"
	"douyin-12306/models"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"sync"
)

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

func (d *UserDAO) Register(ctx context.Context, username string, password string, name string) (*models.User, error) {
	var user models.User

	err := R.Redis.Get(ctx, models.User{}.UsernameKeyPrefix()+username).Err()
	if !errors.Is(err, redis.Nil) {
		if err == nil {
			logger.L.Debug("repository.Register found user in Redis", nil)
			err = errors.New("用户名已存在")
			return nil, err
		}
		// 用日志打印替代返回 error
		//	Redis不可用时的降级策略
		logger.L.Error("Redis Get Error In repository.Register", map[string]interface{}{
			"error": err,
		})
	}

	tx := R.MySQL.Table(models.User{}.TableName()).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check：该username是否已存在
	err = tx.Where(&models.User{Username: username}).Take(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			err = R.Redis.Set(ctx, user.Key(), &user, user.Expiration()).Err()
			if err != nil {
				logger.L.Error("Redis Set Error In userDAO.Register", map[string]interface{}{
					"error": err,
				})
			}
			err = R.Redis.Set(ctx, user.UsernameKey(), user.Id, user.Expiration()).Err()
			if err != nil {
				logger.L.Error("Redis Set Error In userDAO.Register", map[string]interface{}{
					"error": err,
				})
			}
			logger.L.Debug("repository.Register found user in MySQL", nil)
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
	user = models.User{
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

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = R.Redis.Set(ctx, user.Key(), &user, user.Expiration()).Err()
	if err != nil {
		logger.L.Error("Redis Set Error In userDAO.Register", map[string]interface{}{
			"error": err,
		})
	}
	err = R.Redis.Set(ctx, user.UsernameKey(), user.Id, user.Expiration()).Err()
	if err != nil {
		logger.L.Error("Redis Set Error In userDAO.Register", map[string]interface{}{
			"error": err,
		})
	}
	return &user, nil
}
