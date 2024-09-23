package apis

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {

	r.POST("register", Register)

	r.POST("login", Login)
}
