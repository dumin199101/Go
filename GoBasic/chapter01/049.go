package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	//循环读取客户端数据
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("客户端:%v已退出\n", conn.RemoteAddr())
			return
		}
		fmt.Print("接收到来自客户端的数据：", string(buffer[:n]))
		//向客户端做出反馈
		conn.Write([]byte("Hello,客户端，我已收到你发来的数据:" + string(buffer[:n])))
	}
}

func main() {
	// TCP--服务端

	// 1.监听客户端连接
	// 2.服务端与客户端建立连接
	// 3.创建Goroutine处理客户端数据
	listener, e := net.Listen("tcp", "0.0.0.0:8888")
	if e != nil {
		fmt.Println("监听错误：", e)
		return
	}

	for {
		//接收客户端连接
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("服务端连接客户端失败")
			continue
		}

		//创建携程处理客户端数据
		go process(conn)
	}

}
