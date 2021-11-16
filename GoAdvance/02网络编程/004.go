package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP客户端-并发版
func main() {

	conn, e := net.Dial("tcp", "127.0.0.1:10096")

	if e != nil {
		fmt.Println("客户端连接服务器错误", e)
		return
	}

	//读取键盘输入
	for {
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')

		if strings.ToUpper(strings.Trim(readString, "\r\n")) == "Q" {
			fmt.Println("退出...")
			return
		}

		if err != nil {
			fmt.Println("ReadString Error", err)
			return
		}

		//向服务器发送数据
		conn.Write([]byte(readString))

		//读取服务端反馈数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if n == 0 {
			fmt.Println("读取服务端返回的数据结束...")
			return
		}
		if err != nil {
			fmt.Println("Read Error", err)
			return
		}
		fmt.Println("【客户端读取到服务器返回数据】", string(buffer[:n]))
	}

}
