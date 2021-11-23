package main

import "fmt"

// 未关闭的管道，如果管道中空了，再取会阻塞
// 关闭的管道，如果管道空了，再取返回零值

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	for value := range ch {
		fmt.Println(value)
	}
	//v1 := <-ch
	//fmt.Println(v1)
	//v2 := <-ch
	//fmt.Println(v2)
	//v3 := <-ch
	//fmt.Println(v3)
	//v4 := <-ch
	//fmt.Println(v4)
	//v5 := <-ch
	//fmt.Println(v5)
}
