package global

import "time"

type Config struct {
	Server
	Mysql
	Redis
}

type Server struct {
	HTTPPort string
}

type Redis struct {
	Addr string
	Pwd  string
	DB   int
}

type Mysql struct {
	User            string
	Password        string
	Host            string
	Port            string
	DbName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}
