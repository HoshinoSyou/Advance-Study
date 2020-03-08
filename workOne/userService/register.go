package userService

import (
	"Advance/workOne/response"
	"Advance/workOne/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

func Register(r *gin.Context) {
	var u User
	var m UserModel
	u.Username = r.PostForm("username")
	u.Password = r.PostForm("password")
	db = sqlConn.Init()
	db.Where("username = ?", u.Username).First(&m)
	if m.ID >= 0 {
		response.Error(r, 202, "用户名已存在喵！")
		return
	}
	m = UserModel{
		Model:    gorm.Model{},
		Username: u.Username,
		Password: u.Password,
	}
	db.Create(&m)
	response.OK(r, "注册成功喵！")
}
