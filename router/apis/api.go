package apis

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {

	r.POST("register", Register)

	r.POST("login", Login)

	r.POST("logout", Logout)

	r.GET("info", Info)

	r.POST("update", Update)
}
