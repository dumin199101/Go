package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 声明变量
	var i int
	// 初始化变量
	i = 10
	// 使用变量
	fmt.Println(i)

	//自动推导变量类型
	var n = 20
	fmt.Println("n=", n)

	//声明变量并初始化变量
	m := 30
	fmt.Println("m=", m)

	//数据类型
	var a int8 = 10
	fmt.Printf("数据类型:%T,占据字节是:%d\n", a, unsafe.Sizeof(a))

	var b = 3.12
	fmt.Printf("数据类型:%T,占据字节是:%d\n", b, unsafe.Sizeof(b))

	var c byte = 'a'
	fmt.Println("c=", c)    //97
	fmt.Printf("c=%c\n", c) //c

	var d bool = false
	fmt.Println(d)
	fmt.Printf("数据类型:%T,占据字节是:%d\n", d, unsafe.Sizeof(d))

	var e = "hello"
	e = "world"
	fmt.Println(e)

	//默认值
	var f int     //0
	var g float64 //0
	var h bool    //false
	var j string  //""
	fmt.Println("f=", f, "g=", g, "h=", h, "j=", j)

}
