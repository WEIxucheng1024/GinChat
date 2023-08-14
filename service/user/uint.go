package user

import "time"

type UserResp struct {
	UUID       string
	UserName   string
	Name       string
	CreateTime time.Time
	Phone      string
	Email      string
	Token      string
}
