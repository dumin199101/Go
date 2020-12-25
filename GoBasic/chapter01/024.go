package main

import "fmt"

type Stopper2 interface {
	stop()
}

type Car2 struct {
	name string
}

func (c *Car2) stop() {
	fmt.Println("汽车制动")
}

func main() {
	//值接收者跟指针接收者实现接口的区别
	//指针类型接收者，只有指针类型结构体变量可以赋值给接口变量
	var stopper Stopper2
	/*var car = Car2{name:"丰田锐志"}
	stopper = car
	stopper.stop()*/

	var car2 = &Car2{name: "奔驰"}
	stopper = car2
	stopper.stop()

}
