package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	//生产者
	go func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			fmt.Println("生产者生产：", i)
			ch <- i * i
		}
		close(ch)
	}(ch)
	//消费者
	go func(ch <-chan int) {
		for value := range ch {
			fmt.Println("消费者消费：", value)
		}
	}(ch)

	time.Sleep(time.Second)
}
