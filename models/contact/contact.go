package contact

import "gorm.io/gorm"

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  string // 谁的关系
	TargetId string // 对应的谁
	TYpe     int    // 对应的类型 1.人&人	2.人&群
	Desc     string
}

func (this *Contact) TableName() string {
	return "contact"
}
