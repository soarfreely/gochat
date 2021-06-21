package common

import (
	error2 "chatroom/common/error"
	"fmt"
	"net"
)

type Schema struct {
	Protocol string
	Ip string
	Port uint
}

const (
	Protocol = "tcp"
	Ip = "0.0.0.0"
	Port = 9999
)

// 连接服务端
func Connect(schema Schema) (conn net.Conn){
	conn, err := net.Dial(schema.Protocol, Socket(schema))
	if err != nil {
		fmt.Printf("网络连接错误。参数:%s, 错误:%s", schema, err)
	}
	return conn
}

// 连接服务端
func Server() (net.Conn) {
	schema := Schema{
		Protocol: "tcp",
		Ip:       "127.0.0.1",
		Port:     9999,
	}

	return Connect(schema)
}

// 拼接socket-套接字
func Socket(schema Schema) (socket string){
	return fmt.Sprintf("%s:%d", schema.Ip, schema.Port)
}


// 发送消息到服务端
func Send(content []byte)  {
	conn := Server()

	_, err := conn.Write(content)
	if err != nil {
		error2.New(500100, err.Error())
	}

	defer conn.Close()

	fmt.Println("【发送内容】:", string(content))
}

// 接收发送消息
func Receive(conn net.Conn) (buf []byte) {
	var content = make([]byte, 8096)
	_, err := conn.Read(content)

	if err != nil {
		error2.New(500102, err.Error())
	}

	return content
}