package auth

import (
	"ginchat1/models/auth"
)

func SignOut(userId, token string) (resp *LoginResponse, code int, err error) {
	loginLog, err := auth.FindToken(userId, token)
	if err != nil {
		code = 500
		return
	}
	if loginLog.UserID != userId {
		code = 401
		return
	}

	db := auth.DeleteToken(userId, token)
	if db.Error != nil {
		err = db.Error
		code = 500
	}
	code = 200
	return
}
