package main

import (
	"fmt"
	"time"
)

func test001() {
	for i := 1; i < 5; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
}

func test002() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发生了异常")
		}
	}()
	var mapping map[int]string
	mapping[1] = "hello"
}

func main() {

	go test001()

	go test002()

	//主线程
	for i := 1; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(time.Second)
	}
}
