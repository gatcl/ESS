package main

import (
	"ess/config"
	"ess/global"
	"ess/route"
)

func main() {
	if global.DB != nil {
		config.CreateTables(global.DB)
	}
	route.InitRoute()
}
