package main

import (
	"fmt"
	"time"
)

var channelFlag = make(chan bool)

func print(s string) {
	for _, value := range s {
		fmt.Printf("%c\n", value)
		time.Sleep(time.Millisecond * 100)
	}
}

func test1() {
	print("hello")
	channelFlag <- true
}

func test2() {
	<-channelFlag
	print("world")
}

func main() {
	/*
		操作	        nil channel	         closed channel	        not-closed non-nil channel
		close	      panic	                 panic	                  成功 close
		写 ch <-	 一直阻塞	                 panic	                  阻塞或成功写入数据
		读 <- ch	 一直阻塞	                 读取对应类型零值	          阻塞或成功读取数据
	*/

	/*
		  通道阻塞场景：
		     无缓冲通道的特点是，发送的数据需要被读取后，发送才会完成，它阻塞场景：
				通道中无数据，但执行读通道。
				通道中无数据，向通道写数据，但无协程读取。
		     有缓存通道的特点是，有缓存时可以向通道中写入数据后直接返回，缓存中有数据时可以从通道中读到数据直接返回，这时有缓存通道是不会阻塞的，它阻塞的场景是：
				通道的缓存无数据，但执行读通道。
				通道的缓存已经占满，向通道写数据，但无协程读。
		  select是执行选择操作的一个结构，它里面有一组case语句，它会执行其中无阻塞的那一个，如果都阻塞了，那就等待其中一个不阻塞，进而继续执行，
		  它有一个default语句，该语句是永远不会阻塞的，我们可以借助它实现无阻塞的操作。
	*/

	go test1()

	go test2()

	time.Sleep(time.Second * 3)

}
