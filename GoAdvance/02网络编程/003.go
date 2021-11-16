package main

import (
	"fmt"
	"net"
	"strings"
)

// TCP-服务端并发版

func main() {

	listener, e := net.Listen("tcp", "127.0.0.1:10096")
	if e != nil {
		fmt.Println("Listener error", e)
		return
	}
	defer listener.Close()

	//启动多个协程监听客户端
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("Accept error", e)
			return
		}
		go read(conn)
	}

}

func read(conn net.Conn) {
	defer conn.Close()
	//循环读取数据
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if n == 0 {
			fmt.Printf("读取客户端:%s的数据结束...\r\n", conn.RemoteAddr())
			return
		}
		if err != nil {
			fmt.Println("Read Error", err)
			return
		}
		fmt.Println("【服务器读取到客户端数据】", string(buffer[:n]))

		// 向客户端反馈数据
		conn.Write([]byte("客户端你好，我已收到你发来的数据:" + strings.ToUpper(string(buffer[:n]))))

	}

}
