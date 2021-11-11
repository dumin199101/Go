package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的Map

var m = sync.Map{}
var waitgroup = sync.WaitGroup{}

func main() {

	for i := 1; i < 20; i++ {
		waitgroup.Add(1)
		go func(i int) {
			defer waitgroup.Done()
			//设置值
			key := strconv.Itoa(i)
			m.Store(key, i)
			//获取值
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
		}(i)
	}

	waitgroup.Wait()

}
