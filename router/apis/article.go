package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poetize_server/models"
	"poetize_server/models/biz"
	"poetize_server/models/request"
	"poetize_server/models/response"
)

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
