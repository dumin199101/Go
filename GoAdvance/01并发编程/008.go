package main

import "fmt"

//  通道关闭后再取值ok=false,返回通道零值

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("向管道写数据：", i)
			ch <- i
		}
		close(ch)
	}()

	//通道关闭后再取值ok=false
	for {
		if value, ok := <-ch; ok == true {
			fmt.Println("从管道中读数据：", value)
		} else {
			fmt.Println("从管道中读完数据,返回管道零值", value)
			break
		}
	}

}
