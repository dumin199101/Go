package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {

	//Goexit:退出当前协程，并调用defer方法，由于不是panic，所以recover返回nil

	go func() {
		defer func() {
			fmt.Println("协程终止....")
			fmt.Println("recover error....", recover())
		}()

		for i := 1; i < 10; i++ {
			if i == 5 {
				runtime.Goexit()
			}
			fmt.Println("协程执行中" + strconv.Itoa(i))
		}
	}()

	for i := 1; i < 10; i++ {
		fmt.Println("主线程执行中" + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 100)
	}

}
