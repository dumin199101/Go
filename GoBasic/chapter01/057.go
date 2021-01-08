package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func run(ctx context.Context) {
	fmt.Println("协程开启")
	defer wg2.Done()
	for {
		time.Sleep(time.Second)
		select {
		//Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("hello,world")
	}
}

func main() {
	// context包的使用
	//Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
	//WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。
	ctx, cancel := context.WithCancel(context.Background())
	wg2.Add(1)
	go run(ctx)
	//3s后主动结束go协程
	time.Sleep(time.Second * 3)
	cancel()
	wg2.Wait()
	fmt.Println("主线程结束")
}
