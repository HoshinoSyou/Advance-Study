package service

import (
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Ballot   int    `json:"ballot"`
}

func Register(r *gin.Context) {
	var u User
	var m UserModel
	u.Username = r.PostForm("username")
	u.Password = r.PostForm("password")
	db = sqlConn.Mysql()
	db.Where("username = ?", u.Username).First(&m)
	if m.ID >= 0 {
		response.Error(r, 202, "注册失败喵！用户名已存在喵！")
	}
	m = UserModel{
		Model:    gorm.Model{},
		Username: u.Username,
		Password: u.Password,
	}
	db.Create(&m)
	b := UserData{
		Username: u.Username,
		Ballot:   3,
	}
	db.Create(&b)
	response.OK(r, "注册成功喵！")
}
