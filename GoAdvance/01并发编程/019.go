package main

import (
	"fmt"
	"sync"
)

//对比018.go,通过Channel实现安全的并发加操作
var channel = make(chan int, 1)
var v int
var waitg sync.WaitGroup

func main() {
	//初始化
	channel <- v
	for i := 0; i < 10000; i++ {
		waitg.Add(1)
		go func() {
			defer waitg.Done()
			v := <-channel
			v++
			channel <- v
		}()
	}

	waitg.Wait()

	fmt.Printf("当前v的值是%d", <-channel)

}
