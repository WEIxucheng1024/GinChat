package auth

import (
	"ginchat1/models"
	"ginchat1/models/auth"
	"ginchat1/service/user"
	"ginchat1/utils"
	"strconv"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"passWord"`
}

func Login(req *LoginRequest) (resp *user.UserResp, code int, err error) {
	findUser := models.GetUserByUserName(req.UserName)
	if findUser.UserName == "" {
		code = 401
		return
	}
	validPassWord := utils.ValidPassword(req.Password, findUser.Salt, findUser.PassWord)
	if !validPassWord {
		code = 401
		return
	}

	token := utils.RandomString(64, 62)

	db := auth.SaveToken(strconv.Itoa(int(findUser.ID)), token)

	if db.Error != nil {
		err = db.Error
		code = 500
	}

	resp = &user.UserResp{
		ID:         strconv.Itoa(int(findUser.ID)),
		UserName:   findUser.UserName,
		Name:       findUser.Name,
		CreateTime: findUser.CreatedAt,
		Phone:      findUser.Phone,
		Token:      token,
	}
	code = 200
	return
}
