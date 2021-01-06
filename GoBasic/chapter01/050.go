package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// TCP--客户端
	// 1.建立与Server端数据连接
	// 2.发送数据给Server端，并接收Server端返回的数据
	// 3.关闭客户端连接
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("客户端连接服务端失败，失败原因：", err)
		return
	}

	//接收键盘输入
	reader := bufio.NewReader(os.Stdin)

	//向客户端写数据
	for {
		s, err := reader.ReadString('\n')
		str := strings.Trim(s, "\r\n")

		if err != nil {
			fmt.Println("数据读取失败:", err)
		}

		//键盘输入quit退出客户端
		if "Q" == strings.ToUpper(str) {
			break
		}

		//向服务端写入数据
		_, err = conn.Write([]byte(str + "\n"))
		if err != nil {
			fmt.Println("客户端向服务端写入数据失败", err)
			return
		}
		//读取服务端发送过来的数据
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("客户端读取服务端返回数据失败", err)
			return
		}
		fmt.Print("收到来自服务端的反馈：", string(buff[:n]))
	}

	//关闭客户端连接
	defer conn.Close()

}
