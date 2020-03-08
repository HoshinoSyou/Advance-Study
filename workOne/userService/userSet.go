package userService

import (
	"Advance/workOne/response"
	"Advance/workOne/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Information struct {
	gorm.Model
	Username     string `json:"username"`
	NickName     string `json:"nick_name"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Address      string `json:"address"`
	Introduction string `json:"introduction"`
}

func SetInformation(s *gin.Context) {
	var i Information
	i = Information{
		Username:     s.PostForm("username"),
		NickName:     s.PostForm("nickname"),
		Sex:          s.PostForm("sex"),
		Age:          s.GetInt("age"),
		Address:      s.PostForm("address"),
		Introduction: s.PostForm("introduction"),
	}
	db = sqlConn.Init()
	db.Create(&i)
	response.OK(s, "个人信息保存成功喵！")
}
