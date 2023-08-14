package auth

import (
	"ginchat1/utils"
	"gorm.io/gorm"
)

type LoginLog struct {
	gorm.Model
	UserUUID string
	Token    string
}

func (this *LoginLog) TableName() string {
	return "login_log"
}

func SaveToken(userUUID, token string) *gorm.DB {
	loginLog := &LoginLog{
		UserUUID: userUUID,
		Token:    token,
	}
	return utils.DB.Create(loginLog)
}

func DeleteToken(userUUID, token string) *gorm.DB {
	loginLog := &LoginLog{
		UserUUID: userUUID,
		Token:    token,
	}
	return utils.DB.Where("user_uuid = ?", loginLog.UserUUID).Where("token = ?", loginLog.Token).Delete(&loginLog)
}

func FindToken(userUUID, token string) (loginLog *LoginLog, err error) {
	db := utils.DB.Where("user_uuid = ?", userUUID).Where("token = ?", token).Find(&loginLog)
	return loginLog, db.Error
}
