package dal

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"douyin-12306/repo"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var D dal

type dal struct {
	MySQL *gorm.DB
	Redis *redis.Client
}

func Init() {
	db, err := repo.NewDB()
	if err != nil {
		logger.L.Panicf("Init DB Error: %s\n", err)
	}
	D.MySQL = db
	logger.L.Infow("Init DB Success", map[string]interface{}{
		"config": config.C.MySQL,
	})

	rds, err := repo.NewRedisClient()
	if err != nil {
		logger.L.Panicf("Init Redis Client Error: %s\n", err)
	}
	D.Redis = rds
	logger.L.Infow("Init Redis Success", map[string]interface{}{
		"config": config.C.Redis,
	})
}
