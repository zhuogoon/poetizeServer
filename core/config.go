package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"poetize_server/config"
	"poetize_server/global"
)

func InitConfig() {
	const ConfigFile = "setting.yaml"

	s := &config.SettingConfig{}
	file, err := os.ReadFile(ConfigFile)
	if err != nil {
		logrus.Error("配置文件读取失败")
		return
	}
	err = yaml.Unmarshal(file, s)
	if err != nil {
		logrus.Error("配置文件解析失败")
		return
	}
	logrus.Info("配置文件 成功")
	global.Config = s
}
