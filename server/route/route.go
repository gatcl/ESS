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

	if err := r.Run(global.CFG.HTTPPort); err != nil {
		os.Exit(0)
	}
}
