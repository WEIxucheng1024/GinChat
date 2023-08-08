package controllers

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags 首页
// @Success 200 {object} string "welcome"
// @Router /test2 [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "okk2!",
	})
}
