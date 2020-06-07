package routers

import (
	"Advance/SummerExam/response"
	"Advance/SummerExam/ws"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateRoom(c *gin.Context) {
	_, err := ws.Websocket(c, c.Writer, c.Request)
	if err != nil {
		msg := "创建房间失败喵！请与管理员联系喵！"
		log.Printf("http协议升级websocket失败喵！错误信息：%v", msg)
		response.Error(c, msg)
	} else {
		i, cs := ws.CreateRoom()
		msg := "创建房间成功喵！房间号：" + strconv.Itoa(i)
		data, _ := json.Marshal(cs)
		response.OKWithData(c, msg, data)
	}
}
