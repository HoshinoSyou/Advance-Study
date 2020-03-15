package userService

import (
	"Advance/workOne/jwt"
	"Advance/workOne/response"
	"Advance/workOne/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(l *gin.Context) {
	var u User
	var m UserModel
	u.Username = l.PostForm("username")
	u.Password = l.PostForm("password")
	db = sqlConn.Init()
	db.Where(u).First(&m)
	if m.ID > 0 {
		token := jwt.Create(u.Username)
		response.OKWithToken(l, "登录成功喵！", token)
	} else {
		response.Error(l, 201, "密码错误喵！")
	}
}
