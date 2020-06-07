package routers

import (
	"Advance/SummerExam/dao"
	"Advance/SummerExam/response"
	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	onceNew := c.PostForm("onceNew")
	if oldPassword == "" {
		msg := "旧密码不能为空喵！"
		response.Error(c, msg)
	}
	if newPassword == "" {
		msg := "新密码不能为空喵！"
		response.Error(c, msg)
	}
	if onceNew != newPassword {
		msg := "两次输入的密码不相同喵！"
		response.Error(c, msg)
	}
	username := c.Param("username")
	res := dao.Verify(username, oldPassword)
	if res {
		result := dao.ChangePassword(username, newPassword)
		if result {
			msg := "修改密码成功喵！"
			response.OK(c, msg)
		} else {
			msg := "修改密码失败喵！"
			response.Error(c, msg)
		}
	} else {
		msg := "旧密码错误喵！"
		response.Error(c, msg)
	}
}
