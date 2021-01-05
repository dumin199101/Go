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

	for i := 1; i < 5; i++ {
		//让出当前goroutine的执行权限，调度器安排其他等待的任务运行
		runtime.Gosched()
		fmt.Println("hello,go routine")
	}
}
