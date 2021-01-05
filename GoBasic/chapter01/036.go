package main

import (
	"fmt"
)

func putNum(intChannel chan int, num int) {
	for i := 1; i <= num; i++ {
		//数据放入intChannel
		intChannel <- i
	}
	close(intChannel)
}

func getNum(intChannel chan int, intChannel2 chan int, boolChannel chan bool) {
	var flag bool
	for {
		num, ok := <-intChannel
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//证明不是素数
				flag = false
				break
			}
		}

		if flag {
			//将数据放入intChannel2
			intChannel2 <- num
		}
	}
	//将此协程标记读取完毕
	fmt.Println("此协程已取不到数据，退出...")
	boolChannel <- true
}

func main() {
	//利用协程实现求1-10000的素数
	//1.将1-10000放入管道中
	intChannel := make(chan int, 1000)
	go putNum(intChannel, 10000)
	//2.开启4个协程，从管道中取出数据，判断是否是素数
	intChannel2 := make(chan int, 2000)
	boolChannel := make(chan bool, 4)
	for i := 1; i <= 4; i++ {
		go getNum(intChannel, intChannel2, boolChannel)
	}

	//3.关闭intChannel2管道
	go func() {
		for i := 1; i <= 4; i++ {
			//如果取不到数据一直阻塞
			<-boolChannel
		}
		close(intChannel2)
	}()

	//4.遍历素数把结果取出,当channel关闭且数据都读完了，再读数据会读到该数据类型的零值，且第二个返回值为false
	for {
		res, ok := <-intChannel2
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")

}
