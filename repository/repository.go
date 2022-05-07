package repository

import (
	"douyin-12306/logger"
	_ "douyin-12306/repository/mysqlDB"
)

func init() {
	logger.L.Info("init repository success", map[string]interface{}{
		"package":  "repository",
		"function": "init",
	})
}
