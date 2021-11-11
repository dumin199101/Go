package main

import (
	"fmt"
	"sync"
	"time"
)

// 声明一个互斥锁
var mutex = sync.Mutex{}

func print(s string) {
	mutex.Lock()
	for _, value := range s {
		fmt.Printf("%c\n", value)
		time.Sleep(300 * time.Millisecond)
	}
	mutex.Unlock()
}

func main() {

	// 顺序输出：利用加锁的方式阻塞，对比007.go管道方式
	go func() {
		print("google")
	}()

	go func() {
		print("baidu")
	}()

	time.Sleep(5 * time.Second)
}
