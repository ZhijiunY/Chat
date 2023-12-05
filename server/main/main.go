package main

import (
	"net"
)

// 處理和客戶端的通訊
func process(conn net.Conn) {
	// 這裡需要延時關閉conn
	defer conn.Close()

}

// func main() {
// 	fmt.Println("服務器在8889端口監聽...")
// 	listen, err := net.Listen("tcp", "0.0.0.0:8889")
// 	defer listen.Close()

// 	if err != nil {
// 		fmt.Println("net.Listen err=", err)
// 		return
// 	}

// 	for {
// 		fmt.Println("等待客戶端來連接服務器......")
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			fmt.Println("listen.Accept err=", err)
// 		}

// 		go process(conn)
// 	}

// }

func main() {

}
