package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poetize_server/models"
	"poetize_server/models/biz"
	"poetize_server/models/request"
	"poetize_server/models/response"
)

// CreateArt  添加文章
// @Summary 添加文章
// @Description 用户进行文章的添加
// @Tags 文章相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body models.Article true "username"
// @Success 200 {object} response.BaseResponse
// @Router /api/art/create [post]
// @Security ApiKeyAuth
func CreateArt(c *gin.Context) {
	req := models.Article{}
	resp := response.BaseResponse{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	err = biz.CreateArt(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	resp.Code = http.StatusOK
	resp.Msg = "创建文章成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return

}

// ArtInfo   获取文章信息
// @Summary 获取文章信息
// @Description 获取这个文章的所有信息
// @Tags 文章相关
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param user body request.ArtInfoReq true "id"
// @Success 200 {object} response.ArticleInfo
// @Router /api/art/artinfo [post]
// @Security ApiKeyAuth
func ArtInfo(c *gin.Context) {
	req := request.ArtInfoReq{}
	resp := response.ArticleInfo{}

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	art, err := biz.GetInfoById(req.Id)
	if err != nil {
		resp.Code = http.StatusOK
		resp.Msg = "获取数据失败，请稍后尝试"
		c.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	resp.Article.UserId = art.UserId
	resp.Article.SortId = art.SortId
	resp.Article.LabelId = art.LabelId
	resp.Article.ArticleCover = art.ArticleCover
	resp.Article.ArticleContent = art.ArticleContent
	resp.Article.ArticleTitle = art.ArticleTitle
	resp.Article.VideoUrl = art.VideoUrl
	resp.Article.ViewCount = art.ViewCount
	resp.Article.LikeCount = art.LikeCount
	resp.Article.ViewStatus = art.ViewStatus
	resp.Article.Tips = art.Tips
	resp.Article.RecommendStatus = art.RecommendStatus
	resp.Article.CommentStatus = art.CommentStatus

	resp.Code = http.StatusOK
	resp.Msg = "获取成功"
	c.AbortWithStatusJSON(http.StatusOK, resp)
	return

}
