package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"poetize_server/config"
)

var (
	Config *config.SettingConfig
	DB     *gorm.DB
	UserId uint
	RDB    *redis.Client
)
var Sign = []byte("zhuogoon")

var Ctx = context.Background()
