package login

import (
	"fmt"
)

func Login() (code int, err error){
	var (
		userId = ""
		userPwd = ""
	)

	fmt.Println("请输入用户账号")
	fmt.Scanln(&userId)

	fmt.Println("请输入用户账号密码")
	fmt.Scanln(&userPwd)

	fmt.Printf("您的账号:%s, 密码:%s", userId, userPwd)

	// 验证账号密码
	return checkPwd(userId, userPwd)
}

// 验证账号密码
func checkPwd(userId string, userPwd string) (code int, err error) {
	return 0, nil
}
