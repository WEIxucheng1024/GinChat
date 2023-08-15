package models

import (
	"ginchat1/models/auth"
	"ginchat1/models/contact"
	"ginchat1/models/group_basic"
	"ginchat1/models/message"
	"ginchat1/utils"
)

func InitModels() error {
	return utils.DB.AutoMigrate(&UserBasic{}, &auth.LoginLog{}, &message.Message{}, &contact.Contact{}, &group_basic.GroupBasic{})
}
