package main

import (
	"fmt"
)

func Go() {
	fmt.Println("Go Execute")
}

func Php() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误发生了,error:", err)
		}
	}()
	num := 10
	num2 := 0
	fmt.Println(num / num2)
}

func Python() {
	fmt.Println("Phthon Execute")
}

func main() {
	Go()
	Php()
	Python()
}
