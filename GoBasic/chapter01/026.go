package main

import "fmt"

func main() {
	//空接口：没有定义任何方法的接口，可以存储任意类型变量
	var a interface{}
	s := "hello,world"
	a = s
	fmt.Printf("s的类型%T,s的值%s\n", a, a)

	//类型断言：x.(T) x：表示类型为interface{}的变量,T：表示断言x可能是的类型。
	//该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

	if v, ok := a.(string); ok {
		fmt.Println("接口变量的值为：", v)
	} else {
		fmt.Println("断言失败")
	}

}
