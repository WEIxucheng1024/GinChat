package controllers

import (
	"ginchat1/service/auth"
	"github.com/gin-gonic/gin"
)

func RenderJSONAuth(c *gin.Context, resp interface{}, code int, err error) {
	if code != 200 {
		c.JSON(code, gin.H{
			"message": err,
			"code":    code,
		})
	} else {
		c.JSON(code, gin.H{
			"data": resp,
		})
	}
}

func AuthToken() {}

func RenderJson(c *gin.Context, resp interface{}, code int, err error) {
	authCode, authErr := auth.AuthUser(c)
	if authCode != 200 {
		c.JSON(authCode, gin.H{
			"message": authErr.Error(),
			"code":    authCode,
		})
		return
	}
	if code != 200 {
		c.JSON(code, gin.H{
			"message": err.Error(),
			"code":    code,
		})
	} else {
		c.JSON(code, gin.H{
			"data": resp,
		})
	}

}
