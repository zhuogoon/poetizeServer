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

// Register 注册接口
// @Summary 注册
// @Description 注册用户
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user body request.UserRequest true "username"
// @Success 200 {object} response.BaseResponse
// @Router /api/user/register [post]
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

// Login 登录接口
// @Summary 登录
// @Description 用户的登录，通过账号密码
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param user body request.UserRequest true "username"
// @Success 200 {object} response.BaseResponse
// @Router /api/user/login [post]
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
		if err == nil {
			resp.Code = http.StatusOK
			resp.Msg = "用户未注册"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
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

// Logout  登出接口
// @Summary 登出
// @Description 登出自己的账号
// @Tags 用户相关
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.BaseResponse
// @Router /api/user/logout [post]
// @Security ApiKeyAuth
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

// Info 获取信息接口
// @Summary 获取信息
// @Description 获取到所需要的用户基本信息
// @Tags 用户相关
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} response.UserInfo
// @Router /api/user/info [get]
// @Security ApiKeyAuth
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

	resp.Code = http.StatusOK
	resp.Msg = "获取成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return
}

// Update 更新用户信息接口
// @Summary 更新信息
// @Description 更新用户的基本信息
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body models.User true "username"
// @Success 200 {object} response.BaseResponse
// @Router /api/user/update [post]
// @Security ApiKeyAuth
func Update(c *gin.Context) {
	req := models.User{}
	resp := response.BaseResponse{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	err = biz.Update(req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "更新失败"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Msg = "更新成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return
}

// ChangePassword  更新用户面膜
// @Summary 更新密码
// @Description 更新用户的密码
// @Tags 用户相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body request.UserPassword true "password"
// @Success 200 {object} response.BaseResponse
// @Router /api/user/changepassword [post]
// @Security ApiKeyAuth
func ChangePassword(c *gin.Context) {
	req := request.UserPassword{}
	resp := response.BaseResponse{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	err = biz.ChangePassword(req.Password)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "修改失败"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
	token := c.GetHeader("Authorization")
	if token != "" {
		global.RDB.Set(global.Ctx, token, "revoked", 0) // 将令牌加入黑名单，设置为永不过期
	}

	resp.Code = http.StatusOK
	resp.Msg = "修改成功，请重新登录"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return

}
