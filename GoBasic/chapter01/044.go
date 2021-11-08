package main

import (
	"fmt"
	"runtime"
	"time"
)

func aa() {
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
		//time.Sleep(10)
	}
}

func bb() {
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
		//time.Sleep(10)
	}
}

func main() {
	//一个逻辑核心，做完一个再做另一个,在同一个线程内，如果没有阻塞，协程不会让出CPU执行权，打印结果顺序执行
	runtime.GOMAXPROCS(1)
	go aa()
	go bb()
	time.Sleep(time.Second)
}
