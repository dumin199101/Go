package main

import (
	"fmt"
	"sync"
	"time"
)

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
	time.Sleep(time.Second * 5)

	fmt.Println("计算1-20的阶乘中...")

	//打印阶乘结果
	for key, value := range resultmap {
		lock.Lock()
		fmt.Println(key, value)
		lock.Unlock()
	}

}
