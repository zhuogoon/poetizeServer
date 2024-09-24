package core

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "poetize_server/docs" // main 文件中导入 docs 包
)

func InitSwag(r *gin.Engine) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
