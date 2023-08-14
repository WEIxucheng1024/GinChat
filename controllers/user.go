package controllers

import (
	"fmt"
	"ginchat1/models"
	"ginchat1/service/user"

	"github.com/gin-gonic/gin"
)

// GetAllUserBasic
// @Summary 所有用户
// @Tags User
// @Param user-uuid header string true "用户id"
// @Param user-token header string true "用户token"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserList [get]
func GetAllUserBasic(c *gin.Context) {
	resp, code, err := user.GetAllUser()
	RenderJson(c, resp, code, err)
}

// GetUserByUserName
// @Summary 根据userName获取用户
// @Tags User
// @Param userName formData string true "账号"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByUserName [post]
func GetUserByUserName(c *gin.Context) {
	resp, code, err := user.GetUserByUserName(c)
	RenderJson(c, resp, code, err)
}

// GetUserByPhone
// @Summary 根据手机号获取用户
// @Tags User
// @Param phone formData string true "手机号"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByPhone [post]
func GetUserByPhone(c *gin.Context) {
	resp, code, err := user.GetUserByPhone(c)
	RenderJson(c, resp, code, err)
}

// GetUserByEmail
// @Summary 根据邮箱获取用户
// @Tags User
// @Param email formData string true "邮箱"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByUEmail [post]
func GetUserByEmail(c *gin.Context) {
	resp, code, err := user.GetUserByEmail(c)
	RenderJson(c, resp, code, err)
}

// CreateUser
// @Summary 新增用户
// @Tags User
// @Param userName body string true "用户名"
// @Param passWord body string true "密码"
// @Param name body string true "名称"
// @Param email body string false "email"
// @Param phone body string false "phone"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	var req *models.UserBasic
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}
	resp, code, err := user.CreateUser(req)
	RenderJSONAuth(c, resp, code, err)
}

// DeleteUser
// @Summary 删除用户
// @Tags User
// @Param userName query string true "用户名"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	code, err := user.DeleteUser(c)
	RenderJson(c, nil, code, err)
}

// UpdateUser
// @Summary 修改用户
// @Tags User
// @Param userName formData string false "账号"
// @Param name formData string false "名称"
// @Param passWord formData string false "密码"
// @Param phone formData string false "手机号"
// @Param email formData string false "邮箱"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	resp, code, err := user.UpdateUser(c)
	RenderJson(c, resp, code, err)
}
