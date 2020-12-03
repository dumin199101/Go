package main

import (
	"fmt"
)

func main() {
	//算术运算符:/法运算符跟运算数数据类型有关
	var a = 10 / 4
	var b float32 = 10.0 / 4
	var c float32 = 10 / 4
	fmt.Println(a) // 2
	fmt.Println(b) // 2.5
	fmt.Println(c) // 2

	//取余运算符计算公式：a%b = a-a/b*b = 10 - 10 / 3 * 3 = 10 -3 * 3 = 1
	var d int = 10 % 3
	fmt.Println(d)

	// ++ 跟 --运算符，++--运算符只能在后边
	e := 1
	e++
	fmt.Println(e)

}
