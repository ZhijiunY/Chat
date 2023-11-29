package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/ZhijiunY/Chat/common/message"
)

// 寫一個函數，完成登入
func login(userId int, userPwd string) (err error) {

	// 1. 連接到服務器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err=", err)
		return
	}
	// 延時關閉
	defer conn.Close()

	// 2. 準備通過 conn 發送消息給服務
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. 創建一個 LoginMes 結構體
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd	
	
	// 4. 將 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5. 把 data 賦給 mes.Data 字段
	mes.Data = string(data)

	// 6. 將 mess 進行序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=", err)
		return
	}
	}

	// 7. 到這個時候 data 就是我們要發送的消息
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	// 發送長度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil{
		fmt.Println("con.write(buf) fail",err)
		return
	}

	fmt.Printf("客戶端發送的訊息長度＝%d 內容=%s", len(dat), string(data))

	// 發送消息本身
	_, err = conn.Write(date)
	if err != nil{
		fmt.Println("conn.Write(buf) fail", err)
		return
	}
	
	// 這裡還需要處理服務器端返回的消息
	tf := &utils.Transfer{
		Conn:conn,
	}
	
	return
}

