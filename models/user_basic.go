package models

import (
	"ginchat1/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	UUID          string `gorm:"primaryKey" json:"uuid"`
	UserName      string `gorm:"uniqueIndex;size:16" json:"userName"` // uniqueIndex为唯一索引，size位长度
	PassWord      string `json:"passWord" json:"passWord"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone"`
	Email         string `valid:"email" json:"email"`
	Identity      string
	ClentIP       string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	SignOutTime   *time.Time
	IsLogOut      bool
	DeviceInfo    string
	Name          string `json:"name"`
	Test          string
	Salt          string
}

func (this *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserBasic() ([]*UserBasic, error) {
	var users []*UserBasic
	db := utils.DB.Find(&users)
	return users, db.Error
}

func GetUser(userName, phone, email string) []*UserBasic {
	user := []*UserBasic{}
	utils.DB.Where("user_name = ?", userName).Or("phone = ?", phone).Or("email = ?", email).Find(&user)

	return user
}

func GetUserByUserName(userName string) *UserBasic {
	var users *UserBasic
	utils.DB.Where("user_name = ?", userName).Find(&users)
	return users
}

func GetUserByPhone(phone string) *UserBasic {
	var users *UserBasic
	utils.DB.Where("phone = ?", phone).Find(&users)
	return users
}

func GetUserByEmail(email string) *UserBasic {
	var users *UserBasic
	utils.DB.Where("email = ?", email).Find(&users)
	return users
}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Where("user_name = ?", user.UserName).Delete(&user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	// 根据user_name找到特定的数据，去修改name和passWord，如果传的值为空，那么调语句时不会修改
	// 比如我接口传值只传user_name和PassWord，nane为空，那么语句为：
	// UPDATE `user_basic` SET `updated_at`='2023-08-03 15:25:58.731',`pass_word`='aaa222' WHERE user_name = 'a2' AND `user_basic`.`deleted_at` IS NULL
	return utils.DB.Model(&user).Where("user_name = ?", user.UserName).Updates(
		UserBasic{
			Name:     user.Name,
			PassWord: user.PassWord,
			Phone:    user.Phone,
			Email:    user.Email,
		})
}
