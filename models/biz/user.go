package biz

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"poetize_server/global"
	"poetize_server/models"
)

// CreatUser 创建注册用户
func CreatUser(username, password string) error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req := &models.User{
		Username:   username,
		Password:   string(newPassword),
		UserStatus: 1,
		UserType:   2,
		Deleted:    0,
	}
	return global.DB.Model(&models.User{}).Create(&req).Error
}

// IsUser 检测用户是否存在
func IsUser(username string) (bool, error) {
	var user models.User
	err := global.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 用户不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 用户存在
}

// IsPassword 查询密码是否正确
func IsPassword(username, password string) error {
	var user models.User
	err := global.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// GetIdByUsername 通过 用户名 获取 id
func GetIdByUsername(username string) (uint, error) {
	var user models.User
	err := global.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GetInfo() (models.User, error) {
	var user models.User
	err := global.DB.Model(&models.User{}).Where("id = ?", global.UserId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func Update(user models.User) error {
	return global.DB.Model(&models.User{}).Where("id = ?", global.UserId).Updates(user).Error
}
