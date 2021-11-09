package main

import (
	"fmt"
)

// 解决deadlock

func main() {
	ch := make(chan int)
	go func(ch chan int) {
		a := <-ch
		fmt.Println("从管道中取出数据", a)
	}(ch)
	ch <- 1
	fmt.Println("程序执行结束")
}
