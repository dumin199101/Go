package main

import "fmt"

func sumAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

func test(a int) {
	a = a + 10
	fmt.Println("a在test方法体内部的值：", a)
}

func test01(a *int) {
	*a = *a + 10
	fmt.Println("a在test01方法体内部的值：", *a)
}

func test02(a int, args ...int) int {
	sum := 0
	for key, value := range args {
		fmt.Println(key, value)
		sum += value
	}
	return a + sum
}

func test03() {
	defer fmt.Println("Hello")
	defer fmt.Println("World")
	fmt.Println("test03函数执行了...")
	defer fmt.Println("Go")
}

func main() {
	fmt.Println("main函数执行")
	//函数
	a := 10
	b := 20
	var (
		result1 int
		result2 int
	)

	result1, result2 = sumAndSub(a, b)

	fmt.Println("sum:", result1)
	fmt.Println("sub:", result2)

	fmt.Println("**********************")

	//地址传递跟值传递
	test(a)
	fmt.Println("a在test方法体外部的值", a)

	fmt.Println("**********************")
	test01(&a)
	fmt.Println("a在test01方法体外部的值", a)

	fmt.Println("***********************")
	//可变参数
	result3 := test02(10, 20, 30, 40)
	fmt.Println("四个数的和是：", result3)

	//匿名函数（一个函数只使用一次）
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("匿名函数调用的结果：", res)

	fmt.Println("***********************")
	//defer栈，函数执行完再执行的操作，先入后出
	test03()

}

func init() {
	fmt.Println("init函数执行")
}
