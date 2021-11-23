package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()
		for i := 1; i < 6; i++ {
			ch <- i
		}
		// 缺少close会发生死锁，原因是主线程一直等待遍历管道，但是管道空了，造成阻塞
		close(ch)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()

	w.Wait()

}
