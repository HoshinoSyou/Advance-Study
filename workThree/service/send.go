package service

func send(m []byte, c *Client) {
	for client := range CS.ClientGroup {
		if client != c {
			client.Send <- m
		}
	}
	return
} //客户端发送消息
