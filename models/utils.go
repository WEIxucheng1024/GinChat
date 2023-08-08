package models

import "ginchat1/utils"

func InitModels() error {
	return utils.DB.AutoMigrate(&UserBasic{}, &LoginLog{})
}
