package main

import (
	"fmt"
)

/**
  函数进阶：
      1.可变参数
      2.匿名函数
      3.闭包：函数外部引用函数内部变量
*/

func test(a string, c ...int) {
	fmt.Printf("%T,%v\n", c, c)
}

func test2(f func()) {
	f()
}

func test3() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func main() {
	// 可变参数
	test("abc", 1, 2, 3, 4, 5, 6)
	// 匿名函数
	//1.匿名函数自调用
	func(s string) {
		fmt.Println(s)
	}("Hello World")
	//2.匿名函数赋值给变量
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	println("----------------------------------")
	//3.匿名函数作为参数
	test2(func() {
		println("Good Study Go")
	})
	//4.匿名函数作为返回值
	add2 := test3()
	fmt.Println(add2(4, 6))

}
