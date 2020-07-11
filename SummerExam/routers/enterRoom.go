package routers

import (
	"Advance/SummerExam/util/response"
	"Advance/SummerExam/util/ws"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func EnterRoom(e *gin.Context) {
	c, err := ws.Websocket(e, e.Writer, e.Request)
	if err != nil {
		msg := "加入房间失败喵！请与管理员联系喵！"
		log.Printf("http协议升级websocket失败喵！错误信息：%v", msg)
		response.Error(e, msg)
	} else {
		cs := e.Param("clients")
		var css ws.Clients
		json.Unmarshal([]byte(cs), &css)
		ws.EnterRoom(c, css)
		msg := "加入房间成功喵！"
		response.OK(e, msg)
	}
}//EnterRoom 加入房间
