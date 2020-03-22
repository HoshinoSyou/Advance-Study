package service

import "github.com/gorilla/websocket"

type Clients struct {
	ClientGroup map[*Client]bool
	Login       chan *Client
	Logout      chan *Client
	SendMessage chan []byte
} //客户端组

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

var CS = Clients{
	ClientGroup: make(map[*Client]bool),
	Login:       make(chan *Client),
	Logout:      make(chan *Client),
	SendMessage: make(chan []byte),
}
