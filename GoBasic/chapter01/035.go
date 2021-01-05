package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu) // 8
	//返回值：上一次GOMAXPROCS的结果，第一次运行，默认值为电脑CPU核心数
	gomaxprocs := runtime.GOMAXPROCS(4)
	gomaxprocs2 := runtime.GOMAXPROCS(4)
	fmt.Println(gomaxprocs)  // 8
	fmt.Println(gomaxprocs2) // 4
}
