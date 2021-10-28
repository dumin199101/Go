package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 1; i < 5; i++ {
			fmt.Println("hello,world")
		}
	}()

	go func() {
		for i := 6; i < 10; i++ {
			fmt.Println("hello,go")
		}
	}()

	for i := 1; i < 5; i++ {
		//让出当前goroutine的执行权限，调度器安排其他等待的任务运行,跟yield线程类似，但是还会竞争，不一定谁抢到CPU执行权
		runtime.Gosched()
		fmt.Println("hello,go routine")
	}
}
