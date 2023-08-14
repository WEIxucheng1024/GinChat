package user

import (
	"errors"
	"fmt"
	"ginchat1/models"
	"ginchat1/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func GetAllUser() (users []*models.UserBasic, code int, err error) {
	users, err = models.GetUserBasic()
	if err != nil {
		code = 500
	} else {
		code = 200
	}
	return
}

func GetUserByUserName(c *gin.Context) (user *models.UserBasic, code int, err error) {
	user = models.GetUserByUserName(c.PostForm("userName"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func GetUserByPhone(c *gin.Context) (users []*models.UserBasic, code int, err error) {
	user := models.GetUserByPhone(c.PostForm("phone"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func GetUserByEmail(c *gin.Context) (users []*models.UserBasic, code int, err error) {
	user := models.GetUserByEmail(c.PostForm("email"))
	if user.UserName == "" {
		code = 400
		err = errors.New("未找到对应成员")
		return
	} else {
		code = 200
	}
	return
}

func CreateUser(user *models.UserBasic) (resp *UserResp, code int, err error) {
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(user.PassWord, user.Salt)
	fmt.Println(user.PassWord)

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		code = 400
		err = errors.New("创建失败，格式有误")
		return
	}

	users := models.GetUser(user.UserName, user.Phone, user.Email)

	for _, i := range users {
		if i.UserName == user.UserName {
			code = 400
			err = errors.New("注册失败，用户名重复")
			return
		}
		if i.Email == user.Email && i.Email != "" {
			code = 400
			err = errors.New("注册失败，邮箱重复")
			return
		}
		if i.Phone == user.Phone && i.Phone != "" {
			code = 400
			err = errors.New("注册失败，手机号重复")
			return
		}
	}

	user.UUID = utils.RandomString(8, 62)

	createUser := models.CreateUser(user)
	if createUser.Error != nil {
		code = 500
		err = errors.New("注册失败")
		return
	} else {
		code = 200
		resp = &UserResp{
			UUID:       user.UUID,
			UserName:   user.UserName,
			Name:       user.Name,
			CreateTime: user.CreatedAt,
			Phone:      user.Phone,
			Email:      user.Email,
		}
	}
	return
}

func DeleteUser(c *gin.Context) (code int, err error) {
	user := &models.UserBasic{
		UserName: c.Query("userName"),
	}
	createUser := models.DeleteUser(user)
	if createUser.Error != nil {
		code = 500
		err = createUser.Error
		return
	} else {
		code = 200
	}
	return
}

func UpdateUser(c *gin.Context) (resp *UserResp, code int, err error) {
	user := &models.UserBasic{
		UserName: c.PostForm("userName"),
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("passWord"),
		Phone:    c.PostForm("phone"),
		Email:    c.PostForm("email"),
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		code = 400
		err = errors.New("修改失败,手机号/邮箱有误")
		fmt.Println(err)
		return
	}

	createUser := models.UpdateUser(user)
	if createUser.Error != nil {
		code = 500
		err = createUser.Error
		return
	} else {
		resp = &UserResp{
			UUID:       user.UUID,
			UserName:   user.UserName,
			Name:       user.Name,
			CreateTime: user.CreatedAt,
			Phone:      user.Phone,
			Email:      user.Email,
		}
		code = 200
	}
	return
}

// 防止跨域站点的伪造请求
var upGrade = websocket.Upgrader{
	// 设置websocket连接时的限制，这里直接return true，没有做限制
	/**
		限制实例：
		// 允许的域名或 IP 地址列表
	        allowedOrigins := map[string]bool{
	            "https://example.com": true,
	            "https://subdomain.example.com": true,
	            // ... 添加更多允许的来源
	        }

	        // 检查请求的来源是否在允许的列表中
	        origin := r.Header.Get("Origin")
	        return allowedOrigins[origin]
		**/
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func TestMes(c *gin.Context) {
	// 将http连接升级为websocket连接
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, ctx *gin.Context) {
	for {
		// 读取客户端消息
		mesType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 返回消息到客户端
		err = ws.WriteMessage(mesType, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func(message string) {
			time.Sleep(100 * time.Millisecond)
			err = utils.Publish(ctx, utils.PubkushKey, string(message))
			if err != nil {
				fmt.Println(err)
				return
			}
		}(string(message))

		go func() {
			mes, err := utils.Subscribe(ctx, utils.PubkushKey)
			if err != nil {
				fmt.Println(err)
				return
			}
			tm := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("[ws][%s]:[%s]\n", tm, mes)
		}()

	}
}
