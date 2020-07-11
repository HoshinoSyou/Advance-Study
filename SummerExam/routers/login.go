package routers

import (
	"Advance/SummerExam/service"
	"Advance/SummerExam/util/jwt"
	"Advance/SummerExam/util/response"
	"github.com/gin-gonic/gin"
)

func Login(l *gin.Context) {
	res := service.Login(l)
	if res {
		msg := "登录成功喵！"
		username := l.PostForm("username")
		jwt.Create(username)
		response.OK(l, msg)
	} else {
		msg := "用户名或密码错误喵！"
		response.Error(l, msg)
	}
}//Login 登录账户
