package routers

import (
	"Advance/SummerExam/dao"
	"Advance/SummerExam/response"
	"github.com/gin-gonic/gin"
)

func Logout(d *gin.Context) {
	username := d.PostForm("username")
	password := d.PostForm("password")
	res := dao.Logout(username, password)
	if res {
		msg := "删除账户成功喵！不要忘记我们呜呜呜。"
		response.OK(d, msg)
	} else {
		msg := "用户名或密码错误喵！请重新输入后在进行注销喵！"
		response.Error(d, msg)
	}
}
