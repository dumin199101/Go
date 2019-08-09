package main

import (
	"fmt"
	"math"
)

/*
  练习目录：
      1.匿名变量
          匿名变量的特点是一个下画线“_”，但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用。
      2.变量作用域
          任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问的。
      3.整型  int8、int16、int32 和 int64， uint8、uint16、uint32 和 uint64
      4.浮点型 float32 float64
      5.复数 complex64 (32 位实数和虚数)，complex128 (64 位实数和虚数)
      6.布尔类型 true false
      7.字符串
          解释字符串
          非解释字符串
          字符串拼接
          多行字符串
      8.类型转换




*/

func GetData() (int, int) {
	return 100, 200
}

//全局变量
var variable = 300

func VariableArea() {
	// 局部变量
	variabale2 := 400
	fmt.Println(variable, variabale2)
}

// 复数 x*y = (ac-bd,bc+ad)
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i

//布尔类型
// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

//字符串
var str = "Hello Java"
var str1 = `Hello \r \n World`
var str2 = str + str1
var str3 = `
     Hello World
     I am study Go
`

//类型转换
func ChangeType() {
	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	i2 := 100
	fmt.Println(int16(i2))
}

func main() {
	v1, _ := GetData()
	_, v2 := GetData()
	fmt.Println(v1, v2)
	VariableArea()
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
	fmt.Println(btoi(true))
	fmt.Println(str, str1, str2, str3)
	ChangeType()

}
