package main

import (
	"fmt"
	"sync"
	"time"
)

//因为map为引用类型，所以即使函数传值调用，参数副本依然指向映射m, 所以多个goroutine并发写同一个映射m， 写过多线程程序的都知道，对于共享变量，资源，并发读写会产生竞争的， 故共享资源遭到破坏

var resultmap = make(map[int]int, 10)
var lock sync.Mutex //声明一个互斥锁

func jiecheng(n int) {
	var res int
	res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将阶乘结果存入map
	lock.Lock()
	resultmap[n] = res //concurrent map writes
	lock.Unlock()
}

func main() {
	//goroutine

	for j := 1; j <= 20; j++ {
		go jiecheng(j)
	}

	//等待5s，让协程执行完毕
	time.Sleep(time.Second * 10)

	fmt.Println("计算1-20的阶乘中...")

	//打印阶乘结果
	for key, value := range resultmap {
		lock.Lock()
		fmt.Println(key, value)
		lock.Unlock()
	}

}
