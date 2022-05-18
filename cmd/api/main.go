package main

import (
	"douyin-12306/cmd/api/router"
	"douyin-12306/cmd/api/rpc"
	"douyin-12306/config"
	"douyin-12306/logger"
	"douyin-12306/pkg/tracer"

	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(config.C.Services.Api.Name)
	rpc.InitRPC()
}

func main() {
	// 初始化日志
	logger.L = logger.NewZapLogger(config.C.Log.Out, config.C.Log.Level)
	defer func() {
		err := logger.L.Sync()
		if err != nil {
			panic(err)
		}
	}()

	Init()

	gin.SetMode(config.C.Services.Api.Mode)
	r := gin.Default()

	// 注册路由
	router.Register(r)

	// 运行
	if err := r.Run(config.C.Services.Api.Addr); err != nil {
		logger.L.Fatal(err)
	}
}
