package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ZhijiunY/Chat/common/message"
)

// 這裡將這些方法關聯到結構體中
type Transfer struct {
	// 分析他應該有哪些字段
	Conn net.Conn
	Buf  [8096]byte // 這時傳輸時，使用緩衝

}

func (This *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 8096)
	fmt.Println("讀取客戶端發送的數據...")
	// conn.Read 在 conn 沒有被關閉的形況下，才會阻塞
	// 如果客戶關閉的 conn 則，就不會組塞
	_, err = This.Conn.Read(This.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read err=", err)
		// 自定義一個錯誤
		// err = errors.New("read pkg header error")
		return
	}

	// 根據 buf[:4] 轉成一個 uint32 類型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(This.Buf[0:4])
	// 根據 pkgLen 讀取消息內容
	n, err := This.Conn.Read(This.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn.Read fail err=", err)
		// 自定義一個錯誤
		// err = errors.New("read pkg body error")
		return
	}

	// 把 pkgLen 反序列化成 -> message.Message
	// 技術就是一層窗戶紙 &mes!!
	err = json.Unmarshal(This.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (This *Transfer) WriterPkg(data []byte) (err error) {

	// 先發送一個長度給對方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(This.Buf[0:4], pkgLen)
	// 發送長度
	n, err := This.Conn.Write(This.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(buf) fail", err)
		return
	}

	// 發送 data 本身
	n, err = This.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
