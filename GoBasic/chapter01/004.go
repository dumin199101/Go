package main

import "fmt"

//同一个包下，变量函数不能重名,局部变量声明必须使用
//标识符：变量名，函数名，常量都使用驼峰法，首字母大写表示公有，首字母小写表示私有
var a int = 10
var UserName string = "lisi"

func main() {
	fmt.Println(a)
	fmt.Println(UserName)
}
