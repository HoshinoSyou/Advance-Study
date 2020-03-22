package service

import (
	"encoding/json"
	"log"
)

func Surf(c *Client) {
	defer func() {
		CS.Logout <- c
		c.Websocket.Close()
	}()
	for {
		_, m, err := c.Websocket.ReadMessage()
		if err != nil {
			log.Printf("读取用户%v(id:%v)的消息失败喵！错误信息：%v", c.Username, c.ID, err)
			CS.Logout <- c
			c.Websocket.Close()
			return
		} else {
			bytes, err := json.Marshal(&Message{
				Id:       c.ID,
				Username: c.Username,
				Message:  string(m),
			})
			if err != nil {
				log.Printf("序列化用户%v(id:%v)的消息失败，发送失败喵！错误信息：%v", c.Username, c.ID, err)
				return
			}
			CS.SendMessage <- bytes
		}
	}
} //客户端接收并查看消息
