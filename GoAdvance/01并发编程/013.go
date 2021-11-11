package main

import "fmt"

// 死锁常见情况

func main() {

	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	num := <-ch
	fmt.Println("读取管道中的数据：", num)

}

// 死锁2
func deadlock2() {
	ch := make(chan int)
	//死锁原因：在这个地方发生阻塞，程序不会往下运行
	num := <-ch
	fmt.Println("读取管道中的数据", num)
	go func() {
		ch <- 100
	}()
}

// 死锁1
func deadlock1() {
	ch := make(chan int)
	ch <- 100
	//死锁原因：在这个地方发生阻塞，程序不会往下运行
	num := <-ch
	fmt.Println("读取管道中的数据", num)
}
