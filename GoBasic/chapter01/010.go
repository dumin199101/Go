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

// 函数作为形参或者返回值
func addV() func(int) int {
	n := 10
	return func(v int) int {
		n = n + v
		return n
	}
}

// 变量作用域
var name = "zhangsan" //全局变量，整个包有效
func printName() {
	name := "lisi" //局部变量，函数内部有效
	fmt.Println(name)
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

	fmt.Println("*************************")

	//闭包
	sumV := addV()

	sum := sumV(1)
	fmt.Println("第1次累加后的值：", sum)
	sum = sumV(2)
	fmt.Println("第2次累加后的值：", sum)
	sum = sumV(3)
	fmt.Println("第3次累加后的值：", sum)

	//变量作用域
	printName()
	fmt.Println(name)

	var i int                 // 局部变量
	for i := 1; i < 10; i++ { //i变量只在代码块内有效
		fmt.Println("内部i=", i)
	}

	fmt.Println("外部：i=", i)

}

func init() {
	fmt.Println("init函数执行")
}
