package apis

import "github.com/gin-gonic/gin"

func Admin(r *gin.RouterGroup) {

	r.POST("reg", AdminReg)

	r.POST("log", AdminLogin)
}

func Router(r *gin.RouterGroup) {

	r.POST("register", Register)

	r.POST("login", Login)

	r.POST("logout", Logout)

	r.GET("info", Info)

	r.POST("update", Update)

	r.POST("changepassword", ChangePassword)
}

func Art(r *gin.RouterGroup) {

	r.POST("create", CreateArt)

	r.POST("artinfo", ArtInfo)
}
