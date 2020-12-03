package main

import "fmt"

/*Go语言变量声明*/
// 标准格式
var a int
var b float32
var c bool
var d string

// 批量格式
var (
	e int
	f string
	g []float32
	h func() bool
	i struct {
		x int
	}
)

// 简短格式:
/*
需要注意的是，简短模式（short variable declaration）有以下限制：
定义变量，同时显式初始化。
不能提供数据类型。
只能用在函数内部。
*/

func main() {
	m, n := 0, 1
	fmt.Println(m, n)
}
