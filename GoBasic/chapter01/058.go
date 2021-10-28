package main

import "fmt"

func main() {
	// new 用来分配内存，主要用于值类型：int,float,struct返回的是指针
	// make 用来分配内存，主要用于引用类型：channel,slice,map
	var num2 *int
	fmt.Println(num2) // nil,代表的是空指针值
	num2 = new(int)   // *int
	*num2 = 30
	fmt.Printf("num2的类型为%T,num2的地址为%v,num2的值为%v,num2这个指针指向的值为%v", num2, &num2, num2, *num2)

}
