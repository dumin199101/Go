package main

import "fmt"

type Stopper interface {
	stop()
}

type Car struct {
	name string
}

func (c Car) stop() {
	fmt.Println("汽车制动")
}

func main() {
	//值接收者跟指针接收者实现接口的区别
	//值类型接收者，结构体变量跟指针类型结构体变量都可以赋值给接口变量
	var stopper Stopper
	var car = Car{name: "丰田锐志"}
	stopper = car
	stopper.stop()

	var car2 = &Car{name: "奔驰"}
	stopper = car2
	stopper.stop()

}
