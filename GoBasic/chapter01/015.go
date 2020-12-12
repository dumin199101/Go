package main

import (
	"fmt"
)

func Go1() {
	fmt.Println("Go Execute")
}

func Php1() {
	defer func() {
		//捕获错误
		err := recover()
		if err != nil {
			fmt.Println("错误发生了,error:", err)
		}
	}()
	//抛出错误
	panic("被零除的错误")
	//panic之后的代码不会被执行
	fmt.Println("这段代码不会被执行")
}

func Python1() {
	fmt.Println("Phthon Execute")
}

func main() {
	Go1()
	Php1()
	Python1()
}
