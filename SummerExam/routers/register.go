package routers

import (
	"Advance/SummerExam/dao"
	"Advance/SummerExam/response"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Context) {
	username := r.PostForm("username")
	password := r.PostForm("password")
	res := dao.Register(username, password)
	if res {
		msg := "注册成功喵！"
		response.OK(r, msg)
	} else {
		msg := "注册失败喵！该用户名已存在喵！"
		response.Error(r, msg)
	}
}
