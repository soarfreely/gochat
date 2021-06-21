package utils

import (
	"chatroom/common"
	chatError "chatroom/common/error"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

// 100*** 一般错误
// 500*** 系统错误
// 200    成功


func Length(data struct{})  {
	packageLen := uint32(len(ToJson(data)))

	var buf []byte
	binary.BigEndian.PutUint32(buf[0:4], packageLen)

	common.Send(buf[:4])
	fmt.Printf("【发送消息】,长度：%d字节", packageLen)
}

// 结构体转为字节切片
func ToJson(data struct{}) (buf []byte) {
	buf, err := json.Marshal(data)
	if err != nil {
		chatError.New(100000, err.Error())
	}

	return buf
}

func toLength()  {
	
}

func ToStruct(buf []byte) (data struct{}){
	json.Unmarshal(buf, &data)
	
	return data
}


