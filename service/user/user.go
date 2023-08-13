package user

import (
	"errors"
	"fmt"
	"ginchat1/models"
	"ginchat1/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

func GetAllUser() (users []*models.UserBasic, code int, err error) {
	users, err = models.GetUserBasic()
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	return
}

func GetUserByUserName(c *gin.Context) (user *models.UserBasic, code int, err error) {
	user = models.GetUserByUserName(c.PostForm("userName"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func GetUserByPhone(c *gin.Context) (users []*models.UserBasic, code int, err error) {
	user := models.GetUserByPhone(c.PostForm("phone"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func GetUserByEmail(c *gin.Context) (users []*models.UserBasic, code int, err error) {
	user := models.GetUserByEmail(c.PostForm("email"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func CreateUser(user *models.UserBasic) (resp *UserResp, code int, err error) {
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(user.PassWord, user.Salt)
	fmt.Println(user.PassWord)

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		code = 400
		err = errors.New("创建失败，格式有误")
		return
	}

	users := models.GetUser(user.UserName, user.Phone, user.Email)

	for _, i := range users {
		if i.UserName == user.UserName {
			code = 400
			err = errors.New("注册失败，用户名重复")
			return
		}
		if i.Email == user.Email && i.Email != "" {
			code = 400
			err = errors.New("注册失败，邮箱重复")
			return
		}
		if i.Phone == user.Phone && i.Phone != "" {
			code = 400
			err = errors.New("注册失败，手机号重复")
			return
		}
	}

	createUser := models.CreateUser(user)
	if createUser.Error != nil {
		code = 500
		err = errors.New("注册失败")
		return
	} else {
		code = 200
		resp = &UserResp{
			ID:         strconv.Itoa(int(user.ID)),
			UserName:   user.UserName,
			Name:       user.Name,
			CreateTime: user.CreatedAt,
			Phone:      user.Phone,
			Email:      user.Email,
		}
	}
	return
}

func DeleteUser(c *gin.Context) (code int, err error) {
	user := &models.UserBasic{
		UserName: c.Query("userName"),
	}
	createUser := models.DeleteUser(user)
	if createUser.Error != nil {
		code = 500
		err = createUser.Error
		return
	} else {
		code = 200
	}
	return
}

func UpdateUser(c *gin.Context) (resp *UserResp, code int, err error) {
	user := &models.UserBasic{
		UserName: c.PostForm("userName"),
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("passWord"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		code = 400
		err = errors.New("修改失败,手机号/邮箱有误")
		fmt.Println(err)
		return
	}

	createUser := models.UpdateUser(user)
	if createUser.Error != nil {
		code = 500
		err = createUser.Error
		return
	} else {
		resp = &UserResp{
			ID:         strconv.Itoa(int(user.ID)),
			UserName:   user.UserName,
			Name:       user.Name,
			CreateTime: user.CreatedAt,
			Phone:      user.Phone,
			Email:      user.Email,
		}
		code = 200
	}
	return
}
