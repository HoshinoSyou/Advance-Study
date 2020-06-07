package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context,msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    102,
		"message": msg})
	return
}

func OKWithData(c *gin.Context,msg string,data []byte)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    102,
		"message": msg,
		"data": data,
	})
	return
}

func Error(c *gin.Context,msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"code":    202,
		"message": msg})
	return
}
