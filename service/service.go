package service

import (
	"douyin-12306/logger"
)

func init() {
	logger.L.Infow("init service success", map[string]interface{}{
		"package":  "service",
		"function": "init",
	})
}
