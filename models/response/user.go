package response

import (
	"time"
)

// UserInfo 用户获取模型
// @Description UserInfo model with GORM fields
type UserInfo struct {
	BaseResponse
	User struct {
		ID           uint      `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Username     string    `json:"username" gorm:"username"`
		PhoneNumber  string    `json:"phone_number" gorm:"phone_number"`
		Email        string    `json:"email" gorm:"email"`
		Gender       int       `json:"gender" gorm:"gender"`
		OpenId       string    `json:"open_id" gorm:"open_id"`
		Avatar       string    `json:"avatar" gorm:"avatar"`
		Admire       string    `json:"admire" gorm:"admire"`
		Subscribe    string    `json:"subscribe" gorm:"subscribe"`
		Introduction string    `json:"introduction" gorm:"introduction"`
		UserType     int       `json:"user_type" gorm:"user_type"`
	} `json:"user"`
}
