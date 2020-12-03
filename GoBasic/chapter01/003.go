package main

import "fmt"

func main() {
	//值类型：int,float,bool,string,数组，结构体
	//引用类型：Channel,Slice切片，Map，函数，接口，指针

	//指针
	var num int = 10
	var ptr *int = &num
	fmt.Printf("num的地址值：%v\n", ptr)
	*ptr = 20
	fmt.Printf("num现在的值是:%v\n", num)

}
