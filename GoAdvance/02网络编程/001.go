package main

import (
	"fmt"
	"net"
)

/**
  MAC地址-链路层：ARP协议（地址解析协议）
  IP地址-网络层：IP协议
  Port端口-传输层：TCP协议、UDP协议
  应用层：FTP协议、HTTP协议、Telnet协议

  TCP客户端             			TCP服务端
                                   net.Listen()
     net.Dial()----阻塞等待客户连接--->Accept()
     Write()-------数据请求--------> Read()
     Read()<-------数据响应---------Write()
     Close()                       Close()
*/

// TCP服务端

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:10072")
	if e != nil {
		fmt.Println("监听失败", e)
		return
	}
	defer listener.Close()

	//监听客户端连接
	conn, e := listener.Accept()
	if e != nil {
		fmt.Println("客户端连接服务端失败", e)
	}
	//读取客户端数据
	buffer := make([]byte, 1024)
	n, e := conn.Read(buffer)
	if e != nil {
		fmt.Println("Read error", e)
		return
	}

	fmt.Println("读取客户端数据：", string(buffer[:n]))

	conn.Write([]byte("谢谢！！！！"))

	//关闭连接
	defer conn.Close()

}
