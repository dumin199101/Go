package main

import (
	"fmt"
	"runtime"
	"time"
)

func aaa() {
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
	}
}

func bbb() {
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
	}
}

func ccc() {
	for i := 1; i < 100; i++ {
		fmt.Println("C:", i)
	}
}

func main() {
	//总结：单CPU单核中线程只能并发，单CPU多核中线程可以并行。
	//多核多线程，可以把多线程分配给不同的核心处理，其他的线程依旧等待，相当于多个线程并行的在执行，而单核多线程只能是并发。
	//单核多线程指的是单核CPU轮流执行多个线程，通过给每个线程分配CPU时间片来实现，只是因为这个时间片非常短（几十毫秒），
	//所以在用户角度上感觉是多个线程同时执行。
	//多个逻辑核心，同时执行，交替获得CPU执行权
	runtime.GOMAXPROCS(2)
	go aaa()
	go bbb()
	go ccc()
	time.Sleep(time.Second)
}
