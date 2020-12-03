package main

import (
	"fmt"
	"strconv"
)

func main() {
	//其它数据类型转换字符串

	//方式1，使用Sprintf
	var str string
	var a int8 = 10
	str = fmt.Sprintf("%d", a)
	fmt.Printf("type:%T,value:%q\n", str, str)

	var b float32 = 2.0
	str = fmt.Sprintf("%f", b)
	fmt.Printf("type:%T,value:%q\n", str, str)

	var c bool = false
	str = fmt.Sprintf("%t", c)
	fmt.Printf("type:%T,value:%q\n", str, str)

	var d byte = 97
	str = fmt.Sprintf("%c", d)
	fmt.Printf("type:%T,value:%q\n", str, str)

	//方式2:使用strconv函数转换
	var num int64 = 99
	var num1 float64 = 20.00
	var b2 bool = true

	str = strconv.FormatInt(num, 10)
	fmt.Printf("type:%T,value:%q\n", str, str)
	//f代表格式
	str = strconv.FormatFloat(num1, 'f', 1, 64)
	fmt.Printf("type:%T,value:%q\n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("type:%T,value:%q\n", str, str)

	//特殊int与str相互转换
	str = strconv.Itoa(22)
	fmt.Printf("type:%T,value:%q\n", str, str)
	// func Atoi(s string) (int, error)
	var num2 int
	num2, _ = strconv.Atoi("11")
	fmt.Printf("type:%T,value:%d\n", num2, num2)

	//字符串转换基本数据类型
	var str1 string = "11"
	var str2 string = "20.00"
	var str3 string = "true"
	var num3 int64
	var num4 float64
	var b3 bool

	num3, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("type:%T,value:%d\n", num3, num3)

	num4, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("type:%T,value:%f\n", num4, num4)

	b3, _ = strconv.ParseBool(str3)
	fmt.Printf("type:%T,value:%t\n", b3, b3)

}
