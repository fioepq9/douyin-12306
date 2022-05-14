package main

import (
	"douyin-12306/config"
	"douyin-12306/logger"
	"douyin-12306/router"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "douyin-12306/repository"
)

func main() {
	defer func() {
		err := logger.L.Sync()
		if err != nil {
			panic(err)
		}
	}()
	gin.SetMode(config.C.Gin.Mode)

	r := gin.Default()
	router.Register(r)

	err := r.Run(fmt.Sprintf("%s:%s",
		config.C.Main.Host,
		config.C.Main.Port))
	if err != nil {
		panic(err)
	}
}
