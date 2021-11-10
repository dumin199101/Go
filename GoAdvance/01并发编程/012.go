package main

import (
	"fmt"
)

// select多路复用:管道可以不关闭，但是注意不要产生死锁。如果多个channel同时ready，则随机选择一个执行

func main() {

	ch := make(chan int, 5)
	ch2 := make(chan int, 5)
	exit := make(chan bool, 2)

	go func(ch chan<- int) {
		for i := 1; i < 10; i++ {
			ch <- i
			fmt.Println("通道1写入", i)
		}
		//此处如果不关闭，那么当管道数据读取完毕后，再读会阻塞，造成死锁
		close(ch)
	}(ch)

	go func(ch chan<- int) {
		for i := 11; i < 20; i++ {
			ch <- i
			fmt.Println("通道2写入", i)
		}
		//此处如果不关闭，那么当管道数据读取完毕后，再读会阻塞，造成死锁
		close(ch)
	}(ch2)

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("通道1读取", v)
			} else {
				exit <- true
				if len(exit) == 2 {
					fmt.Println("数据读取完毕")
					return
				}
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Println("通道2读取", v)
			} else {
				exit <- true
				if len(exit) == 2 {
					fmt.Println("数据读取完毕")
					return
				}
			}
		}
	}
}
