package main

import "fmt"

func add() func(i int) int {
	var n = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func main() {
	//闭包小案例
	f := add()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
