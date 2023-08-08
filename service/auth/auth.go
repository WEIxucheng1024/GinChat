package auth

import (
	"ginchat1/models"
	"ginchat1/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := models.GetUserByUserName(c.PostForm("userName"))
	if user.UserName == "" {
		c.JSON(400, gin.H{
			"message": "用户名不存在",
		})
		return
	}
	validPassWord := utils.ValidPassword(c.PostForm("passWord"), user.Salt, user.PassWord)
	if !validPassWord {
		c.JSON(400, gin.H{
			"message": "用户名/密码错误",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "登录成功",
		"data":    user,
	})
}
