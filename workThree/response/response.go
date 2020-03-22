package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Normal(n *gin.Context) {
	n.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "服务器正常运行"})
}

func Error(e *gin.Context) {
	e.JSON(200, gin.H{
		"status":  "error",
		"message": "服务器发生故障喵！"})
} //JSON响应
