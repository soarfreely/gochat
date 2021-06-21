package error

import "fmt"

type ChatError struct {
	Code uint
	Endpoint string
	Message string
}

// 结构体,字段首字母一定要大写，否则跨包实例化异常。

func (e *ChatError) Error() (ChatError string) {
	return fmt.Sprintf("chatRoom，错误消息，错误码:%d，错误信息:%s", e.Code, e.Message)
}

func New(code uint, text string) (err *ChatError) {
	return &ChatError{
		Code:    code,
		Message: text,
	}
}