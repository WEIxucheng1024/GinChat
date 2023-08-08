package controllers

import (
	"fmt"
	"ginchat1/models"
	"ginchat1/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
)

// GetAllUserBasic
// @Summary 所有用户
// @Tags User
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserList [get]
func GetAllUserBasic(c *gin.Context) {
	users := models.GetUserBasic()
	c.JSON(200, gin.H{
		"message": users,
	})
}

// GetUserByUserName
// @Summary 根据userName获取用户
// @Tags User
// @Param userName formData string true "账号"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByUserName [post]
func GetUserByUserName(c *gin.Context) {
	user := models.GetUserByUserName(c.PostForm("userName"))
	if user.UserName == "" {
		c.JSON(400, gin.H{
			"message": "未找到对应成员",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": user,
	})
}

// GetUserByPhone
// @Summary 根据手机号获取用户
// @Tags User
// @Param phone formData string true "手机号"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByPhone [post]
func GetUserByPhone(c *gin.Context) {
	user := models.GetUserByPhone(c.PostForm("phone"))
	if user.UserName == "" {
		c.JSON(400, gin.H{
			"message": "未找到对应成员",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": user,
	})
}

// GetUserByEmail
// @Summary 根据邮箱获取用户
// @Tags User
// @Param email formData string true "邮箱"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/getUserByUEmail [post]
func GetUserByEmail(c *gin.Context) {
	user := models.GetUserByEmail(c.PostForm("email"))
	if user.UserName == "" {
		c.JSON(400, gin.H{
			"message": "未找到对应成员",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": user,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags User
// @Param userName formData string true "用户名"
// @Param passWord formData string true "密码"
// @Param name formData string true "名称"
// @Param email formData string false "email"
// @Param phone formData string false "phone"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := &models.UserBasic{
		UserName: c.PostForm("userName"),
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Phone:    c.PostForm("phone"),
	}

	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(c.PostForm("passWord"), user.Salt)
	fmt.Println(user.PassWord)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "创建失败，格式有误",
		})
		return
	}

	users := models.GetUser(user.UserName, user.Phone, user.Email)

	for _, i := range users {
		if i.UserName == user.UserName {
			c.JSON(400, gin.H{
				"message": "注册失败，用户名重复",
			})
			return
		}
		if i.Email == user.Email && i.Email != "" {
			c.JSON(400, gin.H{
				"message": "注册失败，邮箱重复",
			})
			return
		}
		if i.Phone == user.Phone && i.Phone != "" {
			c.JSON(400, gin.H{
				"message": "注册失败，手机号重复",
			})
			return
		}
	}

	createUser := models.CreateUser(user)
	if createUser.Error != nil {
		c.JSON(400, gin.H{
			"message": "注册失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags User
// @Param userName query string true "用户名"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := &models.UserBasic{
		UserName: c.Query("userName"),
	}
	createUser := models.DeleteUser(user)
	if createUser.Error != nil {
		c.JSON(200, gin.H{
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
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
	user := &models.UserBasic{
		UserName: c.PostForm("userName"),
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("passWord"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "修改失败,手机号/邮箱有误",
		})
		return
	}

	createUser := models.UpdateUser(user)
	if createUser.Error != nil {
		c.JSON(200, gin.H{
			"message": "修改失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "修改成功",
	})
}
