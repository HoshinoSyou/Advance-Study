package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(o *gin.Context, msg string) {
	o.JSON(http.StatusOK, gin.H{
		"code":    101,
		"message": msg})
}

func OKWithToken(t *gin.Context, msg string, token string) {
	t.JSON(http.StatusOK, gin.H{
		"code":    102,
		"message": msg,
		"token":   token})
}

func OKWithData(d *gin.Context, msg string, data interface{}) {
	d.JSON(http.StatusOK, gin.H{
		"code":    103,
		"message": msg,
		"data":    data})
}

func Error(e *gin.Context, code int, msg string) {
	e.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg})
}
