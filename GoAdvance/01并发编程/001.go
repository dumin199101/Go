package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	cpunum := runtime.NumCPU()

	fmt.Println(cpunum)

	go func() {
		for i := 1; i < 100; i++ {
			runtime.Gosched()
			fmt.Println("go routine:", i)
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			if i == 50 {
				runtime.Goexit()
			}
			fmt.Println("test routine:", i)
		}
	}()

	for i := 1; i < 100; i++ {
		fmt.Println("main routine:", i)
	}

	time.Sleep(time.Second)
}
