package service

import (
	"Advance/workTwo/response"
	"Advance/workTwo/sqlConn"
	"github.com/gin-gonic/gin"
)

func Exit(e *gin.Context)  {
	rdb = sqlConn.Redis()
	username := e.Param("username")
	rdb.HDel(username)
	response.OK(e,"退出比赛成功喵！")
}
