package biz

import (
	"golang.org/x/crypto/bcrypt"
	"poetize_server/global"
	"poetize_server/models"
)

func CreateAdmin(username, password string) error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req := &models.User{
		Username:   username,
		Password:   string(newPassword),
		UserStatus: 1,
		UserType:   1,
		Deleted:    0,
	}
	return global.DB.Model(&models.User{}).Create(&req).Error
}

func UserType(id uint) (int, error) {
	user := &models.User{}
	err := global.DB.Model(models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return 2, nil
	}
	return user.UserType, nil
}
