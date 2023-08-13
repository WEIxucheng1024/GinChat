package main

import (
	"fmt"
	"ginchat1/models"
	"ginchat1/router"
	"ginchat1/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	err := models.InitModels()
	if err != nil {
		fmt.Println("AutoMigrate err :", err)
		return
	}
	r := router.Router()

	r.Run(":8081")
}
