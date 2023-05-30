package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RC  *redis.Client
	CFG Config

	DateFormat = "2006-01-02 15:04:05"
)
