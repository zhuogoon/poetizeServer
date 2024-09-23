package core

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"poetize_server/global"
)

func InitRedis() {
	addr := global.Config.RedisConfig.GetAddr()
	pwd := global.Config.RedisConfig.GetPwd()

	global.RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
	})

	if err := global.RDB.Ping(global.Ctx).Err(); err != nil {
		logrus.Error("redis连接失败")
		return
	}
	logrus.Info("redis连接成功")
}
