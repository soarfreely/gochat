package main

import (
	error2 "chatroom/common/error"
	"errors"
	"fmt"
	"log"
	"net"
)

// 服务端监听的端口
const (
	port  = 9999
	network = "tcp"
	ip = "0.0.0.0"
	)

func process(conn net.Conn )  {
	var content = make([]byte, 8096)
	conn.Read(content)

	fmt.Println(string(content))
}


func main() {
	// 监听地址
	address := fmt.Sprintf("%s:%d", ip, port)

	// 提示信息
	log.Println("log,正在服务器监听端口:", port)

	// 开启监听
	listener, err := net.Listen("", address)
	if err != nil {
		//fmt.Println("net.Listen,错误信息:", err)
		log.Fatal("net.Listen", &error2.ChatError{
			Code:    500000,
			Message: err.Error(),
		})
	}

	abc := errors.New("error")
	abc.s

	// 等待客户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept,错误信息:", err)
		}

		defer conn.Close()

		// 开启 协程
		go process(conn)
	}

}
