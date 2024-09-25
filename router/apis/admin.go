package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poetize_server/middleware"
	"poetize_server/models/biz"
	"poetize_server/models/request"
	"poetize_server/models/response"
)

// AdminReg  管理员注册
// @Summary 管理员注册
// @Description 管理员进行账号的注册
// @Tags 管理员相关
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body request.UserRequest true "user_id"
// @Success 200 {object} response.UserInfo
// @Router /api/admin/reg [post]
func AdminReg(c *gin.Context) {
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
	if b {
		resp.Code = http.StatusOK
		resp.Msg = "用户名已存在"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	err = biz.CreateAdmin(req.Username, req.Password)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Msg = "注册成功，账号类型为管理员用户"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return
}

// AdminLogin   管理员登录
// @Summary 管理员登录
// @Description 管理员后台管理系统的等录
// @Tags 管理员相关
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body request.UserRequest true "user_id"
// @Success 200 {object} response.UserInfo
// @Router /api/admin/log [post]
func AdminLogin(c *gin.Context) {
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
		resp.Msg = "用户不存在"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	id, err := biz.GetIdByUsername(req.Username)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		return
	}

	userType, err := biz.UserType(id)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	if userType != 1 {
		resp.Code = http.StatusOK
		resp.Msg = "没有权限"
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
