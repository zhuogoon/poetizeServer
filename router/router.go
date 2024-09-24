package router

import (
	"github.com/gin-gonic/gin"
	"poetize_server/router/apis"
)

func Router(r *gin.Engine) {

	user := r.Group("/api/user")
	art := r.Group("api/art")
	admin := r.Group("api/admin")

	apis.Router(user)
	apis.Art(art)
	apis.Admin(admin)
}
