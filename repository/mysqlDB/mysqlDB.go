package mysqlDB

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL mysqlDB

type mysqlDB struct {
	db *gorm.DB
}

func init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C.Mysql.User,
		config.C.Mysql.Passwd,
		config.C.Mysql.Host,
		config.C.Mysql.Port,
		config.C.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.L.Panic(err.Error(), map[string]interface{}{
			"package":      "mysqlDB",
			"function":     "init",
			"mysql config": config.C.Mysql,
		})
	}
	MySQL.db = db
	logger.L.Info("mysql init success", map[string]interface{}{
		"package":      "mysqlDB",
		"function":     "init",
		"mysql config": config.C.Mysql,
	})
}
