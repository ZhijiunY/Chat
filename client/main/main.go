package main

import "fmt"

var (
	userId   int
	userPwd  string
	userName string
)

func main() {

	var key int

	for true {
		fmt.Println("Welcome to the chat system")
		fmt.Println("\t\t\t 1. Log in to the chat room")
		fmt.Println("\t\t\t 2. registered user")
		fmt.Println("\t\t\t 3. Exit system")
		fmt.Println("\t\t\t Please select (1~3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
		case 2:
		case 3:
		default:

		}
	}

	if key == 1 {

		fmt.Println("請輸入用戶的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("請輸入用戶的密碼")
		fmt.Scanf("%s\n", &userPwd)

		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登入失敗")
		} else {
			fmt.Println("登入成功")
		}
	} else if key == 2 {
		fmt.Println("進行用戶註冊")
	}
}
