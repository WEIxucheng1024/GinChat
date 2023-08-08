package service

import "ginchat1/models"

type LoginResponse struct {
	models.UserBasic
	Token string
}
