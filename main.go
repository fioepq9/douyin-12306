package main

import (
	"douyin-12306/config"
	"douyin-12306/router"
	"fmt"
	"github.com/gin-gonic/gin"

	_ "douyin-12306/logger"

	_ "douyin-12306/repository"
)

func main() {
	r := gin.Default()

	gin.SetMode(config.C.Gin.Mode)

	router.Register(r)

	r.Run(fmt.Sprintf("%s:%s",
		config.C.Main.Host,
		config.C.Main.Port))
}
