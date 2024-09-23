package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poetize_server/middleware"
	"poetize_server/models/biz"
	"poetize_server/models/request"
	"poetize_server/models/response"
)

func Register(c *gin.Context) {
	req := &request.UserRequest{}
	resp := &response.BaseResponse{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	err = biz.CreatUser(req.Username, req.Password)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Msg = "注册成功，账号类型为普通用户"
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	req := &request.UserRequest{}
	resp := &response.BaseResponse{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	b, err := biz.IsUser(req.Username)
	if b == false {
		if err != nil {
			resp.Code = http.StatusOK
			resp.Msg = "用户未注册"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		resp.Code = http.StatusOK
		resp.Msg = "服务错误，请稍后再试"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	err = biz.IsPassword(req.Username, req.Password)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "账号或密码错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	jwt, err := middleware.Jwt(req.Username)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "构造token失败"
		return
	}

	c.Header("x-jwt-token", jwt)

	resp.Code = http.StatusOK
	resp.Msg = "登录成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
}
