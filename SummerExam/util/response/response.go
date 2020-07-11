package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context,msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    101,
		"message": msg})
	return
}//返回请求正常，代码101

func OKWithData(c *gin.Context,msg string,data []byte)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    102,
		"message": msg,
		"data": data,
	})
	return
}//返回带有数据的正常请求，代码102

func Error(c *gin.Context,msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    201,
		"message": msg})
	return
}//返回请求异常或错误，代码201
