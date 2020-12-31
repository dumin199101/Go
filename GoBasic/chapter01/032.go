package main

import "fmt"

func main() {
	// channel 管道
	var channelInt = make(chan int, 3)
	//向管道写入
	num := 100
	channelInt <- 10
	channelInt <- num
	channelInt <- 200

	//从管道取出
	num2 := <-channelInt
	fmt.Printf("从管道取出数据：%d\n", num2)

	//管道数据遍历(不关闭管道，all goroutines are asleep - deadlock)
	//关闭管道
	close(channelInt)
	for value := range channelInt {
		fmt.Println(value)
	}

}
