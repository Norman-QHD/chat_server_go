package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	//map锁
	mapLock sync.RWMutex
	//消息广播的channel
	Message chan string
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (server *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Address + "]" + user.Name + ":" + msg

	server.Message <- sendMsg
}

// ListenMessage 监听Message广播消息的channel的goroutine, 一旦有消息就发送给全部的在线user
func (server *Server) ListenMessage() {
	for {
		//服务端收到的广播消息.
		msg := <-server.Message

		//将msg发送给全部的在线user
		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.Channel <- msg
		}
		server.mapLock.Unlock()
	}
}

func (server *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接建立成功: ", conn)

	//创建一个用户对象
	user := NewUser(conn)
	//用户上线了,将用户加入到OnlineMap中.
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	//广播当前用户上线消息
	server.Broadcast(user, "已上线")

	//当前handler阻塞.
	select {}
}

// Start 启动服务器的接口
func (server *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("Net.Listen err: ", err)
		return
	}

	//close listen socket
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Close server error", err)
		}
	}(listener)

	//启动监听message的goroutine
	go server.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accept err: ", err)
			continue
		}

		//do handler
		go server.Handler(conn)
	}
}
