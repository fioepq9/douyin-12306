package repository

import (
	"context"
	"douyin-12306/logger"
	"douyin-12306/models"
	"douyin-12306/responses"
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

// Register
/*	@param
 *
**/
func (d *UserDAO) Register(ctx context.Context, username string, password string, name string) (*models.User, error) {
	var user models.User

	err := R.Redis.Get(ctx, models.User{}.UsernameKeyPrefix()+username).Err()
	if !errors.Is(err, redis.Nil) {
		if err == nil {
			logger.L.Debug("Found User In Redis")
			err = errors.New("用户名已存在")
			return nil, err
		}
		// 用日志打印替代返回 error
		//	Redis不可用时的降级策略
		logger.L.Errorf("Redis Get Error: %s", err.Error())
	}

	tx := R.MySQL.WithContext(ctx).Table(models.User{}.TableName()).Begin()
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
				logger.L.Errorf("Redis Set Error: %s", err.Error())
			}
			err = R.Redis.Set(ctx, user.UsernameKey(), user.Id, user.Expiration()).Err()
			if err != nil {
				logger.L.Errorf("Redis Set Error: %s", err.Error())
			}
			logger.L.Debug("Found User In MySQL")
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
		return nil, err
	}

	err = R.Redis.Set(ctx, user.Key(), &user, user.Expiration()).Err()
	if err != nil {
		logger.L.Errorf("Redis Set Error: %s", err.Error())
	}
	err = R.Redis.Set(ctx, user.UsernameKey(), user.Id, user.Expiration()).Err()
	if err != nil {
		logger.L.Errorf("Redis Set Error: %s", err.Error())
	}
	return &user, nil
}

// GetUserByUsername 根据用户名查询用户
func (d *UserDAO) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// 使用用户名查询用户
	user := models.User{}
	err := R.MySQL.WithContext(ctx).Where("username=?", username).First(&user).Error

	if err != nil {
		// 若此用户名的用户不存在
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserRespByUserId 根据用户id查询用户信息（包含IsFollow信息）
func (d *UserDAO) GetUserRespByUserId(ctx context.Context, selectUserId int64, curUserId int64) (*responses.User, error) {
	var userResp responses.User
	db := R.MySQL

	// 使用连表查询
	err := db.WithContext(ctx).
		Table(models.User{}.TableName()).
		Joins("LEFT JOIN user_follow uf ON `user`.id = uf.user_id AND uf.follower_id=?", curUserId).
		Where(models.User{Id: selectUserId}).
		Select("`user`.*, IF(uf.follower_id IS NULL,FALSE, TRUE) 'is_follow'").
		Take(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &userResp, nil
}
