package service

import (
	"douyin-12306/logger"
)

func init() {
	logger.L.Info("init service success", map[string]interface{}{
		"package":  "service",
		"function": "init",
	})
}
