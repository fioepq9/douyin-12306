package middleware

import (
	"douyin-12306/dto"
	"douyin-12306/logger"
	"douyin-12306/models"
	"douyin-12306/pkg/util"
	"douyin-12306/repository"
	"douyin-12306/responses"
	"douyin-12306/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

// 拒绝访问后面的接口
func notAllow(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
		StatusCode: 2,
		StatusMsg:  "没有用户权限",
	})
}

// CheckTokenAndSaveUser 检查token，如果有则：
// 1.使用token从redis中查询用户信息，保存到上下文中
// 2.刷新token缓存时间
func CheckTokenAndSaveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从query中取出token
		token, hasToken := c.GetQuery("token")
		logger.L.Debugw("CheckTokenAndSaveUser", map[string]interface{}{
			"token in Query": token,
		})
		// 请求中token不存在，直接返回
		if !hasToken {
			c.Next()
			return
		}
		// 2.根据token查询redis，得到userDTO
		userDTO := &dto.UserSimpleDTO{}
		userKey := service.TokenPrefix + token
		err := repository.R.Redis.Get(c, userKey).Scan(userDTO)
		// token在redis中不存在，返回
		if err == redis.Nil {
			logger.L.Debug("token not found in Redis")
			c.Next()
			return
		}

		// 3.将userDTO存储到上下文中
		util.SetUser(c, userDTO)

		// 4.刷新token存在时间
		repository.R.Redis.Expire(c, userKey, models.User{}.Expiration())
		c.Next()
	}
}

// Authorization 对未登录用户进行拦截
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.经过CheckTokenAndSaveUser中间件后查看是否存在用户信息
		user := util.GetUser(c)
		if user == nil {
			// 不存在登录用户，拦截
			notAllow(c)
			return
		}
		// 登录用户存在，放行
		c.Next()
	}
}
