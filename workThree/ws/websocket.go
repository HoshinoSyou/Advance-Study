package ws

import (
	"Advance/workThree/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var w = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    nil,
	Error:           nil,
	CheckOrigin:     nil,
}

func Websocket(c *gin.Context, response http.ResponseWriter, request *http.Request) (err error) {
	conn, err := w.Upgrade(response, request, nil)
	if err != nil {
		log.Printf("升级协议失败喵！错误信息：%v", err)
		return
	}
	defer conn.Close()
	service.Create(c, conn)
	return err
} //将http协议升级为websocket协议
