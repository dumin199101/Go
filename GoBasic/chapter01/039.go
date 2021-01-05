package main

import (
	"fmt"
	"strconv"
)

func main() {
	// select 解决管道数据阻塞问题

	intChannel := make(chan int, 10)
	strChannel := make(chan string, 10)

	for i := 1; i < 10; i++ {
		intChannel <- i
	}

	for i := 1; i < 10; i++ {
		strChannel <- "hello" + strconv.Itoa(i)
	}

	for {
		select {
		case v := <-intChannel:
			fmt.Printf("%d\n", v)
		case v := <-strChannel:
			fmt.Printf("%s\n", v)
		default:
			fmt.Printf("取不到数据了")
			return
		}
	}

}
