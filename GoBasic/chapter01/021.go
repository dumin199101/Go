package main

import "fmt"

type Sayer interface {
	//定义需要实现的方法
	say()
}

type Dogger struct{}

type Catter struct{}

func (d Dogger) say() {
	fmt.Println("汪汪汪")
}

func (c Catter) say() {
	fmt.Println("喵喵喵")
}

func main() {
	//接口的使用
	//接口变量可以存储实现了该接口的实例
	var sayer Sayer
	dog := Dogger{}
	sayer = dog
	sayer.say()
	cat := Catter{}
	sayer = cat
	sayer.say()

}
