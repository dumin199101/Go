package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	user string
	pwd  string
	port int
	host string
)

func main() {
	//获取命令行参数
	args := os.Args
	fmt.Println(args[0]) //go.exe
	fmt.Println(args[1])
	fmt.Println(args[2])

	//通过flag包获取命令行参数
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "123456", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "主机")
	flag.IntVar(&port, "P", 3306, "端口")

	//解析命令行参数
	flag.Parse()

	fmt.Printf("获取的主机：%s,用户名：%s,密码：%s,端口：%d", host, user, pwd, port)

}
