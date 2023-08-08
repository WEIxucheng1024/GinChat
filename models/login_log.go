package models

import (
	"ginchat1/utils"
	"gorm.io/gorm"
)

type LoginLog struct {
	gorm.Model
	UserID string
	Token  string
}

func (this *LoginLog) TableName() string {
	return "login_log"
}

func SaveToken(userId, token string) *gorm.DB {
	loginLog := &LoginLog{
		UserID: userId,
		Token:  token,
	}
	return utils.DB.Create(loginLog)
}
