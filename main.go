package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"poetize_server/core"
	_ "poetize_server/docs" // 引入生成的文档
	"poetize_server/middleware"
	"poetize_server/router"
)

// @title 最美的博客
// @version 1.0
// @description 做一个最没的博客后端
// @host 8080
// @BasePath /api
func main() {
	// 初始化配置文件
	core.InitConfig()

	// 初始化数据库连接
	core.InitDatabase()

	// 初始化redis连接
	core.InitRedis()

	// 初始化路由
	r := gin.Default()
	// 初始化 swag
	core.InitSwag(r)

	// 解决跨域问题
	r.Use(middleware.Cors())

	// jwt权限认证
	r.Use(middleware.JwtParse())

	router.Router(r)

	// 启动项目 在 8080 端口上
	err := r.Run()
	if err != nil {
		logrus.Error("启动失败")
		return
	}
	logrus.Info("服务启动成功")
}
