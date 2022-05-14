package util

import (
	"bytes"
	"douyin-12306/dto"
	"github.com/gin-gonic/gin"
)

func Int64ToString(num int64) string {
	if num == 0 {
		return "0"
	}
	var (
		buf        bytes.Buffer
		numBuf     [19]byte
		i          = 19
		isNegative = num < 0
	)
	if isNegative {
		num = -num
		buf.WriteByte('-')
	}
	for num > 0 {
		i--
		next := num / 10
		numBuf[i] = byte('0' + num - next*10)
		num = next
	}
	for ; i < 19; i++ {
		buf.WriteByte(numBuf[i])
	}
	return buf.String()
}

// SetUser 将用户信息存储到上下文中
func SetUser(c *gin.Context, userDTO *dto.UserSimpleDTO) {
	c.Set("user", userDTO)
}

// GetUser 从上下文中取出用户信息
func GetUser(c *gin.Context) *dto.UserSimpleDTO {
	get := c.MustGet("user")
	return get.(*dto.UserSimpleDTO)
}
