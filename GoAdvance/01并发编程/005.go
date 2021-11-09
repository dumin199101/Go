package main

import (
	"fmt"
	"time"
)

// 主线程一直阻塞，会产生死锁
//此程序不会产生deadlock

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("执行goroutine")
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("程序执行结束")
}
