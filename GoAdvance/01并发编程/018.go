package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全
var l sync.Mutex
var xx int64
var wgg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	xx++ // 等价于上面的操作
	wgg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	xx++
	l.Unlock()
	wgg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&xx, 1)
	wgg.Done()
}

func main() {

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wgg.Add(1)
		//go add()       // 普通版add函数 不是并发安全的
		//go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wgg.Wait()
	end := time.Now()
	fmt.Println(xx)
	fmt.Println(end.Sub(start))
}
