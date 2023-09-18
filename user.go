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
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddr,
		Address: userAddr,
		Channel: make(chan string),
		conn:    conn,
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
