package route

import (
	"ess/global"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	WordGroup(r.Group("word"))
	// 使用中间件
	// WordGroup(r.Group("word", middleware.UserAuth()))

	if err := r.Run(global.CFG.HTTPPort); err != nil {
		os.Exit(0)
	}
}
