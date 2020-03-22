package service

import (
	"github.com/gorilla/websocket"
)

func Input(c *Client) {
	defer c.Websocket.Close()
	for {
		select {
		case m, ok := <-c.Send:
			if ok {
				c.Websocket.WriteMessage(websocket.TextMessage, m)
			} else {
				c.Websocket.WriteMessage(websocket.CloseMessage, []byte{})
			}
		}
	}
} //读取并发送输入的消息
