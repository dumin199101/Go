package main

import "fmt"

/**
  函数
  指针
  结构体
  接口
*/

func sumNum(a int, b int) int {
	return a + b
}

//结构体
type Human struct {
	name  string
	age   int
	addr  string
	adult bool
}

//接口
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	s := 0
	s = sumNum(1, 2)
	fmt.Println(s)

	//指针
	a := 1
	fmt.Println("a:", &a)
	var c *int
	c = &a
	fmt.Println("c:", c)
	//打印指针的值
	fmt.Println("c:", *c)

	//结构体
	var p Human
	p.name = "小明"
	p.age = 22
	p.addr = "北京"
	p.adult = true

	fmt.Println(p.name)

	//接口
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()

}
