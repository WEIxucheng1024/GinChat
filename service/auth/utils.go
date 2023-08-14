package auth

import (
	"errors"
	"ginchat1/models"
	"ginchat1/models/auth"
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	*models.UserBasic
	Token string
}

func AuthUser(c *gin.Context) (code int, err error) {
	userUUID, token := c.GetHeader("user-uuid"), c.GetHeader("user-token")
	if userUUID == "" || token == "" {
		code = 401
		err = errors.New("先随便写一下，这里是鉴权失败，请输入正确id和token")
		return
	}
	loginLog, err := auth.FindToken(userUUID, token)
	if err != nil {
		code = 500
		return
	}
	if loginLog.UserUUID != userUUID {
		code = 401
		err = errors.New("先随便写一下，这里是鉴权失败，请输入正确id和token")
		return
	}
	code = 200
	return
}
