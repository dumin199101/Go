package main

import "fmt"

func writeData(channelInt chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		channelInt <- i
		fmt.Println("写入数据", i)
	}
	//写完就关闭管道，不关会发生死锁
	close(channelInt)
}

func readData(channelInt chan int, channelBool chan bool) {
	for {
		v, ok := <-channelInt
		if !ok {
			break
		}
		fmt.Println("读取数据", v)
	}
	//将标识置为true
	channelBool <- true
	//写完就关闭管道
	close(channelBool)
}

func main() {
	// goroutine 与 channel联合使用
	// 主线程（通过channel控制结束） + 协程（通过channel进行通信）
	// 注意：如果只向管道写入数据，而没有读数据，管道会堵塞，读写频率不一致，这个无所谓。

	//创建两个管道
	channelInt := make(chan int, 50)
	channelBool := make(chan bool, 1)

	//开启协程
	go writeData(channelInt)
	go readData(channelInt, channelBool)

	for {
		//如果channel关闭，且数据都已读取完毕，再读，第二个参数就会返回false
		v, ok := <-channelBool
		if !ok {
			break
		}
		fmt.Printf("读取到结束标识%t,程序运行结束\n", v)
	}
}
