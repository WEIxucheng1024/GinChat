package controllers

import (
	"ginchat1/service/auth"
	"github.com/gin-gonic/gin"
)

// LoginUser
// @Summary 登录用户
// @Tags User
// @Param userName formData string true "账号"
// @Param passWord formData string true "密码"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/loginUser [post]
func LoginUser(c *gin.Context) {
	auth.Login(c)
}
