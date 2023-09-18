package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (server *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接建立成功: ", conn)
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
