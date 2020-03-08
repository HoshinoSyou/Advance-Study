package userService

import (
	"Advance/workOne/response"
	"Advance/workOne/sqlConn"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UploadInformation(u *gin.Context) {
	var i Information
	var db = sqlConn.Init()
	db.Where("username = ?", u.Param("username")).First(&i)
	db.Model(&i).Update(Information{
		Model:        gorm.Model{},
		Username:     u.Param("username"),
		NickName:     u.PostForm("nickname"),
		Sex:          u.PostForm("sex"),
		Age:          u.GetInt("age"),
		Address:      u.PostForm("address"),
		Introduction: u.PostForm("introduction"),
	})
	response.OK(u, "修改信息成功喵！")
}
