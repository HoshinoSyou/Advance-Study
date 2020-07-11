package service

import (
	"Advance/SummerExam/dao"
	"github.com/gin-gonic/gin"
)

//从前端获取的数据与数据库dao层交互 账号部分

func Login(l *gin.Context) bool {
	username := l.PostForm("username")
	password := l.PostForm("password")
	res := dao.Login(username, password)
	return res
}//登录账户

func Register(r *gin.Context) bool {
	username := r.PostForm("username")
	password := r.PostForm("password")
	res := dao.Register(username, password)
	return res
}//注册账户

func Logout(d *gin.Context) bool {
	username := d.PostForm("username")
	password := d.PostForm("password")
	res := dao.Logout(username, password)
	return res
}//注销账户

func ChangePassword(c *gin.Context) bool {
	username := c.PostForm("username")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	res := dao.ChangePassword(username, oldPassword, newPassword)
	return res
}//修改密码
