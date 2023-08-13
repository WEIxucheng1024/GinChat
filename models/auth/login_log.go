package auth

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

func DeleteToken(userId, token string) *gorm.DB {
	loginLog := &LoginLog{
		UserID: userId,
		Token:  token,
	}
	return utils.DB.Where("user_id = ?", loginLog.UserID).Where("token = ?", loginLog.Token).Delete(&loginLog)
}

func FindToken(userId, token string) (loginLog *LoginLog, err error) {
	db := utils.DB.Where("user_id = ?", userId).Where("token = ?", token).Find(&loginLog)
	return loginLog, db.Error
}
