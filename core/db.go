package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"poetize_server/global"
)

func InitDatabase() {
	if global.Config.MysqlConfig.Host == "" {
		logrus.Warn("未配置Mysql，取消连接")
		return
	}
	dsn := global.Config.MysqlConfig.Dsn()

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
		return
	}
	global.DB = db
}
