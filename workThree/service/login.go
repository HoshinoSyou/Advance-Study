package service

import (
	"encoding/json"
	"log"
)

func Login(c *Client) {
	CS.ClientGroup[c] = true
	bytes, err := json.Marshal(&Message{
		Id:       c.ID,
		Username: c.Username,
		Message:  "欢迎加入聊天室喵！"})
	if err != nil {
		log.Printf("用户%v(ID:%v)加入聊天室信息序列化失败喵！错误信息：%v", c.Username, c.ID, err)
		return
	}
	send(bytes, c)
	return
} //客户端接入客户端组内
