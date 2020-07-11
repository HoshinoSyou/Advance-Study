package middleware

import (
	"Advance/SummerExam/util/jwt"
	"Advance/SummerExam/util/response"
	"github.com/gin-gonic/gin"
)

//Token 作为中间价，用于鉴权

func Token(t *gin.Context) {
	header := t.GetHeader("Authorization")
	if len(header) < 7 {
		msg := "没有权限！"
		response.Error(t, msg)
		t.Abort()
		return
	}
	token := header[7:]
	username, err := jwt.CheckToken(token)
	if err != nil {
	msg :=	"没有权限！"
		response.Error(t, msg)
		t.Request.URL.Path = "/userAccount/login/:username/:password"
		t.Abort()
		return
	}
	t.Set("username", username)
	t.Next()
	return
}
