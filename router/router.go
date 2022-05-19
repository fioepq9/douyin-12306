package router

import (
	"douyin-12306/config"
	"douyin-12306/controller"
	"douyin-12306/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// 超时中间件
	apiRouter.Use(middleware.Timeout(time.Duration(config.C.Gin.Timeout) * time.Second))
	// basic apis
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	// 检查token并设置用户信息中间件
	apiRouter.Use(middleware.CheckTokenAndSaveUser())
	apiRouter.GET("/feed/", controller.Feed)

	// 来自extra apis - II，它不需要登录认证
	apiRouter.GET("/comment/list/", controller.CommentList)

	// 为部分接口添加登录拦截中间件
	apiRouter.Use(middleware.Authorization())
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
