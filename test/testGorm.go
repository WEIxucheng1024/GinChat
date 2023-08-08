package main

import (
	"fmt"
	"ginchat1/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic("Something went wrong!")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "测试1"
	//db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // 根据整型主键查找

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("PassWord", "33333")
	db.Model(user).Update("LoginTime", time.Now())
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(user, 1)
}
