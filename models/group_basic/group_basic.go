package group_basic

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId string
	Icon    string // 群头像？
	Desc    string // 描述
	Type    int    // 群类型
}

func (this *GroupBasic) TableName() string {
	return "group_basic"
}
