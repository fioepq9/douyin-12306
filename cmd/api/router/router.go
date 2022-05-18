package router

import (
	"douyin-12306/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// 超时中间件
	// apiRouter.Use(middleware.Timeout(time.Duration(config.C.Gin.Timeout) * time.Second))
	// basic apis
	// apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", handlers.RegisterUser)
	apiRouter.POST("/user/login/", handlers.LoginUser)

	// 登录验证中间件
	// apiRouter.Use(middleware.Authorization())
	apiRouter.GET("/user/", handlers.QueryUserInfo)
	// apiRouter.POST("/publish/action/", controller.Publish)
	// apiRouter.GET("/publish/list/", controller.PublishList)

	// // extra apis - I
	// apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	// apiRouter.GET("/favorite/list/", controller.FavoriteList)
	// apiRouter.POST("/comment/action/", controller.CommentAction)
	// apiRouter.GET("/comment/list/", controller.CommentList)

	// // extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction)
	// apiRouter.GET("/relation/follow/list/", controller.FollowList)
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
