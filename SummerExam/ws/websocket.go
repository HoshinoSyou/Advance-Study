package ws

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Clients struct {
	ClientGroup map[*Client]bool
	Login       chan *Client
	Logout      chan *Client
	SendMessage chan []byte
}

type Client struct {
	ID        string
	Username  string
	Websocket *websocket.Conn
	Send      chan []byte
} //客户端

type Message struct {
	Id       string `json:"ID"`
	Username string `json:"Username"`
	Message  string `json:"message"`
} //消息

var i = 0

func Websocket(c *gin.Context, response http.ResponseWriter, request *http.Request) (client *Client, err error) {
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Printf("升级协议失败喵！错误信息：%v", err)
		return
	}
	defer conn.Close()
	client = Create(c, conn)
	return client, err
}

func Create(c *gin.Context, conn *websocket.Conn) *Client {
	cl := &Client{
		uuid.NewV4().String(),
		c.Param("Username"),
		conn,
		make(chan []byte)}
	return cl
}

func CreateRoom() (int, Clients) {
	var ClientSlice []Clients
	ClientSlice[i] = Clients{
		ClientGroup: make(map[*Client]bool),
		Login:       make(chan *Client),
		Logout:      make(chan *Client),
		SendMessage: make(chan []byte),
	}
	cs := ClientSlice[i]
	i++
	return i, cs
}

func EnterRoom(c *Client, cs Clients) {
	cs.ClientGroup[c] = true
	m := Message{
		Id:       c.ID,
		Username: c.Username,
		Message:  "欢迎进入房间喵！"}
	bytes, err := json.Marshal(&m)
	if err != nil {
		log.Printf("用户%v(ID:%v)加入房间信息序列化失败喵！错误信息：%v", c.Username, c.ID, err)
		send([]byte("欢迎用户"+m.Username+"("+m.Id+")"+"加入房间喵！"), c, cs)
		return
	}
	send(bytes, c, cs)
	return
}

func ExitRoom(c *Client, cs Clients) {
	if cs.ClientGroup[c] == true {
		cs.ClientGroup[c] = false
		delete(cs.ClientGroup, c)
		close(c.Send)
		m := Message{
			Id:       c.ID,
			Username: c.Username,
			Message:  "用户已退出房间喵！"}
		bytes, err := json.Marshal(&m)
		if err != nil {
			log.Printf("用户%v(ID:%v)退出聊天室失败喵！错误信息：%v", c.Username, c.ID, err)
			send([]byte("用户"+m.Username+"("+m.Id+")"+"已退出房间喵！"), c, cs)
			return
		}
		send(bytes, c, cs)
		return
	}
}

func send(m []byte, c *Client, cs Clients) {
	for client := range cs.ClientGroup {
		if client != c {
			client.Send <- m
		}
	}
	return
}

func Surf(c *Client,cs Clients) {
	defer func() {
		cs.Logout <- c
		c.Websocket.Close()
	}()
	for {
		_, m, err := c.Websocket.ReadMessage()
		if err != nil {
			log.Printf("读取用户%v(id:%v)的消息失败喵！错误信息：%v", c.Username, c.ID, err)
			cs.Logout <- c
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
			cs.SendMessage <- bytes
		}
	}
}
