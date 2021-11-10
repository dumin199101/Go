package main

import (
	"fmt"
	"time"
)

// Timer定时器
/**
  type Timer struct{
      C <-chan Time
  }

*/

func main() {
	fmt.Println(time.Now())
	time1 := time.NewTimer(time.Second)
	fmt.Println(<-time1.C)
	fmt.Println(<-time.After(time.Second))

	time2 := time.NewTicker(time.Second)

	for i := 1; i < 10; i++ {
		if i == 5 {
			//重置定时器，2s打印一次
			time2.Reset(2 * time.Second)
		}
		if i == 8 {
			//停止计时器
			time2.Stop()
			break
		}
		<-time2.C
		fmt.Println("每秒钟打印一次，循环打印第", i)
	}

}
