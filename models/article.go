package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserId          uint   `json:"user_id" gorm:"user_id"`
	SortId          int    `json:"sort_id" gorm:"sort_id"`
	LabelId         int    `json:"label_id" gorm:"label_id"`
	ArticleCover    string `json:"article_cover" gorm:"article_cover"`
	ArticleTitle    string `json:"article_title" gorm:"article_title"`
	ArticleContent  string `json:"article_content" gorm:"article_content"`
	VideoUrl        string `json:"video_url" gorm:"video_url"`
	ViewCount       int    `json:"view_count" gorm:"view_count"`
	LikeCount       int    `json:"like_count" gorm:"like_count"`
	ViewStatus      int    `json:"view_status" gorm:"view_status"`
	Password        string `json:"password" gorm:"password"`
	Tips            string `json:"tips" gorm:"tips"`
	RecommendStatus int    `json:"recommend_status" gorm:"recommend_status"` // 是否推荐 0 否， 1 是
	CommentStatus   int    `json:"comment_status" gorm:"comment_status"`
	Deleted         int    `json:"deleted" gorm:"deleted"`
}
