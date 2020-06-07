package routers

import (
	"Advance/SummerExam/response"
	"Advance/SummerExam/ws"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func ExitRoom(e *gin.Context)  {
	c, err := ws.Websocket(e, e.Writer, e.Request)
	if err != nil {
		msg := "退出房间失败喵！请与管理员联系喵！"
		log.Printf("http协议升级websocket失败喵！错误信息：%v", msg)
		response.Error(e, msg)
	} else {
		cs := e.Param("clients")
		var css ws.Clients
		json.Unmarshal([]byte(cs), &css)
		ws.ExitRoom(c, css)
		msg := "退出房间成功喵！"
		response.OK(e, msg)
	}
}
