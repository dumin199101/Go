package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
  读写锁在Go语言中使用sync包中的RWMutex类型。

  读写锁分为两种：读锁和写锁。
        当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
        当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

  两种方式比较：93ms 1.3s

*/

// 读写锁与互斥锁的比较
var wg = sync.WaitGroup{}
var lock = sync.Mutex{}
var rwlock = sync.RWMutex{}
var x int64 //定义全局变量，模拟共享数据

func read() {

	//lock.Lock()
	rwlock.RLock()
	//模拟读请求，耗时1毫秒
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()
	defer wg.Done()
}

func write() {

	//lock.Lock()
	rwlock.Lock()
	//模拟写任务，耗时10毫秒
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	//lock.Unlock()
	rwlock.Unlock()

	defer wg.Done()

}

func main() {

	start := time.Now()

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}
