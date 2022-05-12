package repository

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		config.C.Mysql.User,
		config.C.Mysql.Passwd,
		config.C.Mysql.Host,
		config.C.Mysql.Port,
		config.C.Mysql.DBName,
	)
	R.MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logger.L.Panic("Init MySQL fail", map[string]interface{}{
			"error":        err,
			"MySQL config": config.C.Mysql,
		})
	}
	logger.L.Info("Init MySQL success", map[string]interface{}{
		"MySQL config": config.C.Mysql,
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
		logger.L.Panic("Init Redis fail", map[string]interface{}{
			"error":        err,
			"Redis config": config.C.Redis,
		})
	}
	R.Redis = redis.NewClient(opt)
	logger.L.Info("Init Redis success", map[string]interface{}{
		"Redis config": config.C.Redis,
	})

	logger.L.Info("Init repository success", map[string]interface{}{
		"package":  "repository",
		"function": "init",
	})
}
