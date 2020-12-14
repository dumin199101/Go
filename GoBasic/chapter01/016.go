package main

import "fmt"

func main() {
	a := new(int)
	//变量a的类型是*int,变量a的值是0xc00000c0b0,变量a的地址是0xc000006028,变量a所指向的值是0
	fmt.Printf("变量a的类型是%T,变量a的值是%v,变量a的地址是%v,变量a所指向的值是%v", a, a, &a, *a)
}
