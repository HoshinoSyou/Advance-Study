package service

func Manage() {
	select {
	case c := <-CS.Login:
		Login(c)
	case c := <-CS.Logout:
		Logout(c)
	case m := <-CS.SendMessage:
		SendMessage(m)
	}
	return
} //客户端注册、注销、群发消息管理
