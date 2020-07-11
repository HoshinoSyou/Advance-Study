package routers

import (
	"Advance/SummerExam/service"
	"Advance/SummerExam/util/response"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Context) {
	res := service.Register(r)
	if res {
		msg := "注册成功喵！"
		response.OK(r, msg)
	} else {
		msg := "注册失败喵！该用户名已存在喵！"
		response.Error(r, msg)
	}
}//Register 注册
