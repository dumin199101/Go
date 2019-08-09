package main

import "fmt"

/*Go语言变量的初始化*/

// 1.标准格式
var aa int = 10

// 2.编译器推导类型格式
var attack = 40
var defence = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate

func main() {
	// 3.简短格式声明
	hp := 1000
	fmt.Println(hp)
	fmt.Println(damage)
}
