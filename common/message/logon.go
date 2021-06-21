package message

import (
	"chatroom/common"
)

const (
	// 请求消息
	reqType = "request"

	// 响应消息
	resType = "response"
)

// 客户端向服务端发送登录请求的消息结构
type ReqMessage struct {
	Type string
	Data common.User
}


