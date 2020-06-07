package middleware

import (
	"Advance/SummerExam/jwt"
	"Advance/SummerExam/response"
	"github.com/gin-gonic/gin"
)

func Token(t *gin.Context) {
	header := t.GetHeader("Authorization")
	if len(header) < 7 {
		msg := "修改别人的信息可不是好事，要登录自己的喵！"
		response.Error(t, msg)
		t.Abort()
		return
	}
	token := header[7:]
	username, err := jwt.CheckToken(token)
	if err != nil {
	msg :=	"修改别人的信息可不是好事，要登录自己的喵！"
		response.Error(t, msg)
		t.Request.URL.Path = "/userAccount/login/:username/:password"
		t.Abort()
		return
	}
	t.Set("username", username)
	t.Next()
	return
}
