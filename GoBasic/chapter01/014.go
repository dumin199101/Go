package main

import (
	"fmt"
)

func Go() {
	fmt.Println("Go Execute")
}

func Php() {

	//异常必须先捕获，再触发
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误发生了,error:", err)
		}
	}()

	var p *int = new(int)
	*p = 20

	num := 10
	num2 := 0
	fmt.Println(num / num2)
	fmt.Println("Hello,我会执行么？")
}

func Python() {
	fmt.Println("Phthon Execute")
}

func main() {
	Go()
	Php()
	Python()
}
