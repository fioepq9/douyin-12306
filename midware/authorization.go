package midware

import (
	"douyin-12306/dto"
	"douyin-12306/models"
	"douyin-12306/pkg/util"
	"douyin-12306/repository"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

// 拒绝访问后面的接口
func notAllow(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "没有用户权限",
	})
	c.Abort()
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从query中取出token
		token, ok := c.GetQuery("token")
		// token不存在
		if !ok {
			notAllow(c)
			return
		}
		// 2.查询redis，得到userDTO
		userDTO := &dto.UserSimpleDTO{}
		userKey := service.TokenPrefix + token
		err := repository.R.Redis.Get(c, userKey).Scan(userDTO)
		// token在redis中不存在
		if err == redis.Nil {
			notAllow(c)
			return
		}

		// 3.将userDTO存储到上下文中
		util.SetUser(c, userDTO)

		// 4.刷新token存在时间
		repository.R.Redis.Expire(c, userKey, models.User{}.Expiration())
	}
}
