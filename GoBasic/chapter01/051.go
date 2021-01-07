package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//定时器基本使用
	/*timer := time.NewTimer(time.Second * 3)
	<-timer.C  //一直阻塞到3s时间到,只能执行一次
	fmt.Println("执行Hello，World")*/

	//实现延时功能
	/*//1.Sleep
	time.Sleep(time.Second)
	//2.
	timer1 := time.NewTimer(time.Second)
	<-timer1.C
	//3.相当于newTimer().C
	<-time.After(time.Second)*/

	//停止定时器
	/*timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("定时器执行hello")
	}()


	stop := timer2.Stop()
	if stop{
		fmt.Println("定时器已停止")
	}*/

	//重置定时器
	/*fmt.Println(time.Now())
	timer4 := time.NewTimer(time.Second * 10)
	timer4.Reset(time.Second * 2)
	fmt.Println(<-timer4.C)*/

	//时间到了多次执行
	ticker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	i := 0
	//add的值跟启动的携程数相同
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				break
			}
		}
	}()

	wg.Wait()
}
