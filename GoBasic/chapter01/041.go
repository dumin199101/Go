package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

// 原理跟java多线程的join线程差不多，等待别的线程运行完了，我再运行

func hello(i int) {
	defer wg.Done() //goroutine执行完计数器-1
	fmt.Println("hello" + strconv.Itoa(i))
}

func main() {
	//利用WaitGroup实现Goroutine同步

	for i := 1; i < 15; i++ {
		//执行一个goroutine，计数器+1
		wg.Add(1) //计数器+1
		go hello(i)
	}
	//等待所有goroutine结束
	wg.Wait() //等待计数器为0

	fmt.Println("主线程结束")

}
