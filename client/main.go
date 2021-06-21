package main

import (
	"chatroom/client/login"
	"chatroom/common"
	"fmt"
	"os"
)

const (
	logon    = "1"
	register = "2"
	exit     = "3"
)

// 默认保持登录状态(不退出)
var isExit = false

func menu()  {
	fmt.Println("=欢迎使用golang聊天软件=\n")
	fmt.Println("=========", logon, "-登录========")
	fmt.Println("=========", register, "-注册========")
	fmt.Println("=========", exit, "-退出========\n")

	fmt.Println("请选择:\n")
}

func choice()  {
	var choice = ""

	fmt.Scanln(&choice)
	// fmt.Print(choice)

	switch choice {
		case logon:
			fmt.Println("【进入登录页面】")
			code, err := login.Login()
			if code != 0 || err != nil {
				fmt.Printf("登录失败，错误码:%d,错误信息:%s", code, err)
			}

			// 发送消息
			common.Send([]byte("login success"))
		case register:
			fmt.Println("注册")
		case exit:
			fmt.Println("退出")
			isExit = true
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
	}
}

func main()  {
	// 输出菜单
	menu()

	choice()

	// 处理输入选项
	//for !isExit {
	//	choice()
	//}

	fmt.Println("ok")
}

