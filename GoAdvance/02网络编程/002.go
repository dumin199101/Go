package main

import (
	"fmt"
	"net"
)

// TCP客户端

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:10072")
	if e != nil {
		fmt.Println("Dial Error", e)
		return
	}
	defer conn.Close()

	// 客户端向服务器发送数据
	conn.Write([]byte("Hello Go!"))

	//客户端读取服务器响应数据
	buffer := make([]byte, 1024)
	n, e := conn.Read(buffer)
	if e != nil {
		fmt.Println("Read Error", e)
	}

	data := string(buffer[:n])

	fmt.Println("读取服务端响应数据：", data)

}
