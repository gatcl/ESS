package config

import (
	"context"
	"ess/global"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func initRedis() {
	rcl := redis.NewClient(&redis.Options{
		Addr:     global.CFG.Addr,
		Password: global.CFG.Pwd,
		DB:       global.CFG.DB,
	})
	_, err := rcl.Ping(context.Background()).Result()
	if err == nil {
		global.RC = rcl
	}
}

func initMysql() {
	m := global.CFG.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	//	设置连接池
	p, _ := db.DB()
	p.SetMaxIdleConns(m.MaxIdleConns)
	p.SetMaxOpenConns(m.MaxOpenConns)
	p.SetConnMaxLifetime(m.ConnMaxLifetime)
	global.DB = db

}

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		os.Exit(0)

	} else {
		fmt.Println("成功创建表结构")
	}

}
