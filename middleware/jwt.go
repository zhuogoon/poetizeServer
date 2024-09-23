package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"poetize_server/global"
	"poetize_server/models/biz"
	"poetize_server/models/response"
	"strings"
	"time"
)

type JwtStruct struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// Jwt 添加jwt
func Jwt(username string) (string, error) {
	jwts := &JwtStruct{}
	jwts.Username = username
	jwts.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 2))
	jwts.IssuedAt = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwts)
	signingString, err := token.SignedString(global.Sign)
	if err != nil {
		return "", err
	}
	return signingString, nil
}

// JwtParse 解析jwt
func JwtParse() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := &response.BaseResponse{}

		if c.Request.URL.Path == "/api/user/login" || c.Request.URL.Path == "/api/user/register" {
			return
		}

		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			resp.Code = http.StatusOK
			resp.Msg = "你没登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}

		// 检查黑名单
		val, err := global.RDB.Get(global.Ctx, auth).Result()
		if err == nil && val == "revoked" {
			resp.Code = http.StatusOK
			resp.Msg = "请重新登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}

		t := strings.Split(auth, " ")
		if len(t) != 2 {
			resp.Code = http.StatusOK
			resp.Msg = "你没登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		token := t[1]
		claims := &JwtStruct{}

		withClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return global.Sign, nil
		})
		if err != nil {
			resp.Code = http.StatusOK
			resp.Msg = "你没登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		if withClaims == nil || !withClaims.Valid || claims.Username == "" {
			resp.Code = http.StatusOK
			resp.Msg = "你没登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		c.Set("username", claims.Username)
		id, err := biz.GetIdByUsername(claims.Username)
		if err != nil {
			resp.Code = http.StatusOK
			resp.Msg = "你没登录"
			c.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
		c.Set("userid", id)
		global.UserId = id
	}
}
