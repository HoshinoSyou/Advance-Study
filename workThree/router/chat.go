package router

import (
	"Advance/workThree/response"
	"Advance/workThree/ws"
	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	err := ws.Websocket(c, c.Writer, c.Request)
	if err != nil {
		response.Normal(c)
	} else {
		response.Error(c)
	}
} //路由
