package auth

import (
	"ginchat1/models/auth"
)

func SignOut(userUUID, token string) (resp *LoginResponse, code int, err error) {
	loginLog, err := auth.FindToken(userUUID, token)
	if err != nil {
		code = 500
		return
	}
	if loginLog.UserUUID != userUUID {
		code = 401
		return
	}

	db := auth.DeleteToken(userUUID, token)
	if db.Error != nil {
		err = db.Error
		code = 500
	}
	code = 200
	return
}
