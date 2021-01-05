package main

import (
	"fmt"
	"runtime"
	"time"
)

func aa() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func bb() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	//一个逻辑核心，做完一个再做另一个
	runtime.GOMAXPROCS(1)
	go aa()
	go bb()
	time.Sleep(time.Second)
}
