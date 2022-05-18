package repo

import (
	"douyin-12306/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	gormOpentracing "gorm.io/plugin/opentracing"
)

func NewDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C.MySQL.User,
		config.C.MySQL.Passwd,
		config.C.MySQL.Host,
		config.C.MySQL.Port,
		config.C.MySQL.DBName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger: NewMySQLLogger(config.C.MySQL.Log.Out, config.C.MySQL.Log.Level, gormLogger.Config{
			SlowThreshold:             config.C.MySQL.Log.SlowThreshold * time.Millisecond,
			IgnoreRecordNotFoundError: config.C.MySQL.Log.IgnoreRecordNotFoundError,
		}),
	})
	if err != nil {
		return nil, err
	}
	if err = db.Use(gormOpentracing.New()); err != nil {
		return nil, err
	}
	return db, nil
}
