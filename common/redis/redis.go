package redis

import (
	"fmt"
	"net"
)

// 连接redis
func Connect() (conn net.Conn) {
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis 连接错误:", err)
	}

	return conn
}
