package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func response(s bool, d interface{}, m string, c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Response{s, d, m},
	)
}

func OK(c *gin.Context) {
	response(true, nil, "success", c)
}

func OKMsg(m string, c *gin.Context) {
	response(true, nil, m, c)
}

func OKData(d interface{}, c *gin.Context) {
	response(true, d, "success", c)
}

func OKMsgData(m string, d interface{}, c *gin.Context) {
	response(true, d, m, c)
}

func Fail(c *gin.Context) {
	response(false, nil, "fail", c)
}

func FailMsg(m string, c *gin.Context) {
	response(false, nil, m, c)
}

func FailData(d interface{}, c *gin.Context) {
	response(false, d, "fail", c)
}

func FailMsgData(m string, d interface{}, c *gin.Context) {
	response(false, d, m, c)
}
