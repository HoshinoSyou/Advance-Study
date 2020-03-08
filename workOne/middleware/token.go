package middleware

import (
	"Advance/workOne/jwt"
	"Advance/workOne/response"
	"github.com/gin-gonic/gin"
)

func Token(t *gin.Context) {
	header := t.GetHeader("Authorization")
	if len(header) < 7 {
		response.Error(t, 301, "偷偷改别人的信息可不是好事，要先登录喵！")
		t.Abort()
		return
	}
	token := header[7:]
	username, err := jwt.CheckToken(token)
	if err != nil {
		response.Error(t, 301, "偷偷改别人的信息可不是好事，要先登录喵！")
		t.Request.URL.Path = "/login/:username/:password"
		t.Abort()
		return
	}
	t.Set("username", username)
	t.Next()
	return
}
