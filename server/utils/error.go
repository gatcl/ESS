package utils

import (
	"errors"
	"strings"
)

func MakeError(args ...interface{}) error {
	c := make([]string, len(args))
	for _, v := range args {
		switch v.(type) {
		case string:
			c = append(c, v.(string))
		case error:
			c = append(c, v.(error).Error())
		default:
			panic("类型错误")
		}
	}
	return errors.New(strings.Join(c, ""))
}
