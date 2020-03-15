package service

import (
	"Advance/workTwo/jwt"
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Login(l *gin.Context) {
	var u User
	var m UserModel
	u.Username = l.PostForm("username")
	u.Password = l.PostForm("password")
	db = sqlConn.Mysql()
	db.Where(u).First(&m)
	if m.ID < 0 {
		token := jwt.Create(u.Username)
		response.OKWithToken(l, "登陆成功喵！", token)
	} else {
		response.Error(l, 201, "登陆失败喵！")
	}
}
