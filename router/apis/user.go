package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poetize_server/global"
	"poetize_server/middleware"
	"poetize_server/models"
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

// Logout 退出登录
func Logout(c *gin.Context) {
	resp := &response.BaseResponse{}
	token := c.GetHeader("Authorization")
	if token != "" {
		global.RDB.Set(global.Ctx, token, "revoked", 0) // 将令牌加入黑名单，设置为永不过期
	}
	resp.Code = http.StatusOK
	resp.Msg = "退出成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func Info(c *gin.Context) {
	resp := &response.UserInfo{}
	var err error
	var user models.User
	user, err = biz.GetInfo()
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "获取信息失败"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	resp.User.Username = user.Username
	resp.User.ID = user.ID
	resp.User.PhoneNumber = user.PhoneNumber
	resp.User.Email = user.Email
	resp.User.Gender = user.Gender
	resp.User.OpenId = user.OpenId
	resp.User.Avatar = user.Avatar
	resp.User.Admire = user.Admire
	resp.User.Subscribe = user.Subscribe
	resp.User.Introduction = user.Introduction
	resp.User.UserType = user.UserType
	resp.User.Introduction = user.Introduction
	resp.User.CreatedAt = user.CreatedAt
	resp.User.UpdatedAt = user.UpdatedAt
	resp.User.DeletedAt = user.DeletedAt

	resp.Code = http.StatusOK
	resp.Msg = "获取成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return
}
