package routers

import (
	"Advance/SummerExam/dao"
	"Advance/SummerExam/response"
	"github.com/gin-gonic/gin"
)

func Login(l *gin.Context) {
	username := l.PostForm("username")
	password := l.PostForm("password")
	res := dao.Login(username, password)
	if res {
		msg := "登录成功喵！"
		response.OK(l, msg)
	} else {
		msg := "用户名或密码错误喵！"
		response.Error(l, msg)
	}
}
