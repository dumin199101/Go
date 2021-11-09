package main

import (
	"fmt"
	"time"
)

// 利用管道进行协程通讯（操作共享数据放入管道），解决线程安全问题（CSP并发模型）

var ch = make(chan int)

func printer(s string) {
	for _, value := range s {
		fmt.Printf("%c\n", value)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// 抢占输出
	//go printer("hello")
	//go printer("world")

	// 利用管道阻塞，顺序输出
	go func() {
		printer("google")
		ch <- 100
	}()

	go func() {
		<-ch
		printer("baidu")
	}()

	time.Sleep(5 * time.Second)
}
