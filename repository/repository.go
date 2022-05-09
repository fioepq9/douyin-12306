package repository

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var R repository

type repository struct {
	MySQL *gorm.DB
}

func init() {
	var err error
	// init mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C.Mysql.User,
		config.C.Mysql.Passwd,
		config.C.Mysql.Host,
		config.C.Mysql.Port,
		config.C.Mysql.DBName)
	R.MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.L.Panic("MySQL initial fail", map[string]interface{}{
			"package":      "repository",
			"function":     "init",
			"error":        err,
			"mysql config": config.C.Mysql,
		})
	}

	//
	logger.L.Info("init repository success", map[string]interface{}{
		"package":  "repository",
		"function": "init",
	})
}
