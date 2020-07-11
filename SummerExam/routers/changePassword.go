package routers

import (
	"Advance/SummerExam/service"
	"Advance/SummerExam/util/response"
	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	res := service.ChangePassword(c)
	if res {
		msg := "修改密码成功喵！"
		response.OK(c, msg)
	} else {
		msg := "用户名或密码错误！修改密码失败喵！"
		response.Error(c, msg)
	}
}//ChangePassword 修改账户密码
