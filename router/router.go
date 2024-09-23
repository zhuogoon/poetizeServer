package router

import (
	"github.com/gin-gonic/gin"
	"poetize_server/router/apis"
)

func Router(r *gin.Engine) {

	user := r.Group("/api/user")

	apis.Router(user)
}
