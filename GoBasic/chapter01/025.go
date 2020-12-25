package main

import "fmt"

type Weaponer interface {
	skill()
}

type Flyer interface {
	fly()
}

type Airplane interface {
	Weaponer
	Flyer
}

type Hongzhaji struct {
	name string
}

func (h Hongzhaji) skill() {
	fmt.Println("轰炸机开枪杀人")
}

func (h Hongzhaji) fly() {
	fmt.Println("轰炸机可以起飞")
}

func main() {
	//接口嵌套
	var airplane Airplane
	airplane = Hongzhaji{name: "轰炸机"}
	airplane.fly()
	airplane.skill()

}
