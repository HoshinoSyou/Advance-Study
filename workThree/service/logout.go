package service

import (
	"encoding/json"
	"log"
)

func Logout(c *Client) {
	if CS.ClientGroup[c] == true {
		CS.ClientGroup[c] = false
		delete(CS.ClientGroup, c)
		close(c.Send)
		bytes, err := json.Marshal(&Message{
			Id:       c.ID,
			Username: c.Username,
			Message:  "用户已退出聊天室喵！"})
		if err != nil {
			log.Printf("用户%v(ID:%v)退出聊天室失败喵！错误信息：%v", c.Username, c.ID, err)
			return
		}
		send(bytes, c)
		return
	}
} //注销客户端
