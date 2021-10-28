package main

import (
	"log"
)

func main() {
	//日志操作

	//设置日志格式
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	log.SetPrefix("[log-login]:")

	//设置日志输出到文件
	//log.SetOutput()

	// 2020/12/30 14:45:21 Hello,日志开始记录
	log.Println("Hello,日志开始记录")
	//Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	log.Panicln("对不起，抛出异常了")
	log.Fatalln("对不起，出现错误")

}
