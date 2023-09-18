package main

import (
	"fmt"
	"net"
)

type User struct {
	Name    string
	Address string
	Channel chan string
	conn    net.Conn
	server  *Server
}

// Online 用户的上线业务
func (user *User) Online() {
	//用户上线了,将用户加入到OnlineMap中.
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.Broadcast(user, "已上线")
}

// Offline 用户的下线业务
func (user *User) Offline() {
	//用户上线了,将用户加入到OnlineMap中.
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.Broadcast(user, "↓下线")
}

// DoMessage 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	user.server.Broadcast(user, msg)
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Address: userAddr,
		Channel: make(chan string),
		conn:    conn,
		server:  server,
	}

	//启动监听当前user channel消息的 goroutine
	go user.ListenMessage()

	return user
}

// ListenMessage 监听当前user channel的方法,用一个go来承载,一旦有消息,就发送给对端的客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.Channel

		write, err := user.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("回写消息失败, write: ", write)
			return
		}
	}
}
