package controllers

import (
	"fmt"
	"ginchat1/service/auth"
	"github.com/gin-gonic/gin"
)

// LoginUser
// @Summary 登录用户
// @Tags User
// @Param userName body string true "账号"
// @Param passWord body string true "密码"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/loginUser [post]
func LoginUser(c *gin.Context) {
	var req *auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		return
	}
	resp, code, err := auth.Login(req)
	RenderJSONAuth(c, resp, code, err)
}

// SignOut
// @Summary 登录登出
// @Tags User
// @Param user-id header string true "用户id"
// @Param user-token header string true "用户token"
// @Success 200 {string} json {"code": "200", "message": "Success"}
// @Router /user/signOut [get]
func SignOut(c *gin.Context) {
	_, code, err := auth.SignOut(c.GetHeader("user-id"), c.GetHeader("user-token"))
	RenderJSONAuth(c, nil, code, err)
}
