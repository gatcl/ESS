package utils

import (
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func UUID() string {
	return uuid.NewV4().String()
}

func Args2JSON(p interface{}, c *gin.Context) bool {
	if err := c.ShouldBindJSON(p); err != nil {
		FailMsg(MakeError("参数错误", err).Error(), c)
		return false
	}
	return true
}
