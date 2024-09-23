package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"username"`
	Password     string `json:"password" gorm:"password"`
	PhoneNumber  string `json:"phone_number" gorm:"phone_number"`
	Email        string `json:"email" gorm:"email"`
	UserStatus   int    `json:"user_status" gorm:"user_status"`
	Gender       int    `json:"gender" gorm:"gender"`
	OpenId       string `json:"open_id" gorm:"open_id"`
	Avatar       string `json:"avatar" gorm:"avatar"`
	Admire       string `json:"admire" gorm:"admire"`
	Subscribe    string `json:"subscribe" gorm:"subscribe"`
	Introduction string `json:"introduction" gorm:"introduction"`
	UserType     int    `json:"user_type" gorm:"user_type"`
	UpdateBy     string `json:"update_by" gorm:"update_by"`
	Deleted      int    `json:"deleted" gorm:"deleted"`
}
