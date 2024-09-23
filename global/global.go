package global

import (
	"gorm.io/gorm"
	"poetize_server/config"
)

var (
	Config *config.SettingConfig
	DB     *gorm.DB
	UserId uint
)
var Sign = []byte("zhuogoon")
