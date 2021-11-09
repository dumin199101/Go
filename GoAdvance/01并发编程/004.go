package main

import (
	"fmt"
	"runtime"
)

/**
	执行它我们会发现以下现象:

	有时会发生抢占式输出(说明Go开了不止一个原生线程，达到了真正的并行)
	有时会顺序输出, 打印完0再打印1, 再打印2(说明Go开一个原生线程，单线程上的goroutine不阻塞不松开CPU)
	那么，我们还会观察到一个现象，无论是抢占地输出还是顺序的输出，都会有那么两个数字表现出这样的现象:

	一个数字的所有输出都会在另一个数字的所有输出之前
	原因是， 3个goroutine分配到至多2个线程上，就会至少两个goroutine分配到同一个线程里，单线程里的goroutine 不阻塞不放开CPU, 也就发生了顺序输出

    0 0 0 0 0 1 1 0 0 1 0 0 1 0 1 2 1 2 1 2 1 2 1 2 1 2 2 2 2 2
	0 0 0 0 0 0 0 0 0 0 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2 2 2
	0 0 0 0 0 0 0 1 1 1 1 1 0 1 0 1 0 1 2 1 2 1 2 2 2 2 2 2 2 2
	0 0 0 0 0 0 0 1 1 1 1 1 1 1 1 1 1 0 2 0 2 0 2 2 2 2 2 2 2 2
	0 0 0 0 0 0 0 1 0 0 1 0 1 2 1 2 1 2 1 2 1 2 1 2 1 2 1 2 2 2

*/

var quit1 chan int = make(chan int)

func loop1(id int) { // id: 该goroutine的标号
	for i := 0; i < 10; i++ { //打印10次该goroutine的标号
		fmt.Printf("%d ", id)
	}
	quit1 <- 0
}

func main() {
	runtime.GOMAXPROCS(2) // 最多同时使用2个核

	for i := 0; i < 3; i++ { //开三个goroutine
		go loop1(i)
	}

	for i := 0; i < 3; i++ {
		<-quit1
	}
}
