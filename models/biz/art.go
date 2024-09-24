package biz

import (
	"poetize_server/global"
	"poetize_server/models"
)

func CreateArt(a *models.Article) error {
	a.UserId = global.UserId
	return global.DB.Model(models.Article{}).Create(&a).Error
}

func GetInfoById(id uint) (models.Article, error) {
	art := models.Article{}

	err := global.DB.Model(models.Article{}).First(&art).Error
	if err != nil {
		return art, err
	}
	return art, nil
}
