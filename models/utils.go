package models

import (
	"ginchat1/models/auth"
	"ginchat1/utils"
)

func InitModels() error {
	return utils.DB.AutoMigrate(&UserBasic{}, &auth.LoginLog{})
}
