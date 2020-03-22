package service

func SendMessage(m []byte) {
	for client := range CS.ClientGroup {
		select {
		case client.Send <- m:
			client.Send <- m
		default:
			close(client.Send)
			delete(CS.ClientGroup, client)
		}
	}
} //向所有客户端发送消息
