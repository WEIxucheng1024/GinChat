package router

import (
	"ginchat1/controllers"
	"ginchat1/docs"
	"ginchat1/service/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/test2", controllers.GetIndex)
	r.GET("/user/getUserList", controllers.GetAllUserBasic)
	r.POST("/user/getUserByUserName", controllers.GetUserByUserName)
	r.POST("/user/getUserByUEmail", controllers.GetUserByEmail)
	r.POST("/user/getUserByPhone", controllers.GetUserByPhone)
	r.POST("/user/createUser", controllers.CreateUser)
	r.GET("/user/deleteUser", controllers.DeleteUser)
	r.POST("/user/updateUser", controllers.UpdateUser)
	r.POST("/user/loginUser", controllers.LoginUser)
	r.GET("/user/signOut", controllers.SignOut)

	r.GET("/user/testMes", user.TestMes)
	return r
}
