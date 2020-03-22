package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

func Create(c *gin.Context, conn *websocket.Conn) {
	cl := &Client{
		uuid.NewV4().String(),
		c.Param("Username"),
		conn,
		make(chan []byte)}
	CS.Login <- cl
	go Surf(cl)
	go Input(cl)
} //创建客户端
