package config

import (
	"ess/global"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

func init() {
	loadCFG()
	initRedis()
	initMysql()
}

func loadCFG() {
	f, e := ini.Load("config/config.ini")
	if e != nil {
		fmt.Println("读取配置文件失败")
	}
	// Server
	global.CFG.HTTPPort = f.Section("Server").Key("HTTPPort").MustString(":20618")
	// Redis
	global.CFG.Addr = f.Section("Redis").Key("Addr").MustString("")
	global.CFG.Pwd = f.Section("Redis").Key("Pwd").MustString("")
	global.CFG.DB = f.Section("Redis").Key("DB").MustInt(0)
	// Mysql
	global.CFG.User = f.Section("Mysql").Key("User").MustString("root")
	global.CFG.Password = f.Section("Mysql").Key("Password").MustString("")
	global.CFG.Host = f.Section("Mysql").Key("Host").MustString("localhost")
	global.CFG.Port = f.Section("Mysql").Key("Port").MustString("3306")
	global.CFG.DbName = f.Section("Mysql").Key("DbName").MustString("")
	global.CFG.MaxIdleConns = f.Section("Mysql").Key("MaxIdleConns").MustInt()
	global.CFG.MaxOpenConns = f.Section("Mysql").Key("MaxOpenConns").MustInt()
	global.CFG.ConnMaxLifetime = f.Section("Mysql").Key("ConnMaxLifetime").MustDuration(10 * time.Second)
}
