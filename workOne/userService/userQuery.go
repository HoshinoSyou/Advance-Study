package userService

import (
	"Advance/workOne/response"
	"Advance/workOne/sqlConn"
	"github.com/gin-gonic/gin"
)

func QueryUser(u *gin.Context) {
	var i Information
	i.Username = u.Param("username")
	db = sqlConn.Init()
	db.Where("username = ?", i.Username).First(&i)
	if i.ID < 0 {
		response.Error(u, 203, "该用户不存在喵！")
		return
	}
	response.OKWithData(u, "您查询的用户信息如下喵！", i)
}
