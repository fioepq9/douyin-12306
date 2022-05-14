package repository

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var R repository

type repository struct {
	MySQL *gorm.DB
	Redis *redis.Client
}

func init() {
	var err error
	// 初始化 MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C.MySQL.User,
		config.C.MySQL.Passwd,
		config.C.MySQL.Host,
		config.C.MySQL.Port,
		config.C.MySQL.DBName,
	)
	R.MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: NewMySQLLogger(config.C.MySQL.Log.Out, config.C.MySQL.Log.Level, gormLogger.Config{
			SlowThreshold:             time.Duration(config.C.MySQL.Log.SlowThreshold) * time.Millisecond,
			IgnoreRecordNotFoundError: config.C.MySQL.Log.IgnoreRecordNotFoundError,
		}),
	})
	if err != nil {
		logger.L.Panicw("Init MySQL fail", map[string]interface{}{
			"error":        err,
			"MySQL config": config.C.MySQL,
		})
	}
	logger.L.Infow("Init MySQL success", map[string]interface{}{
		"MySQL config": config.C.MySQL,
	})
	// 初始化 Redis
	redisURL := fmt.Sprintf("redis://%s:%s@%s/%s",
		config.C.Redis.Username,
		config.C.Redis.Password,
		config.C.Redis.Addr,
		config.C.Redis.DB,
	)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		logger.L.Panicw("Init Redis fail", map[string]interface{}{
			"error":        err,
			"Redis config": config.C.Redis,
		})
	}
	R.Redis = redis.NewClient(opt)
	logger.L.Infow("Init Redis success", map[string]interface{}{
		"Redis config": config.C.Redis,
	})

	logger.L.Info("Init repository success")
}
