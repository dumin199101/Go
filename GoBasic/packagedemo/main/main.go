package main

import (
	"Go/GoBasic/packagedemo/utils"
	"fmt"
)

func main() {
	// 引入其他包的变量，默认从GoPath的src目录开始寻址，调用方式：包名.变量名
	fmt.Println(utils.DbUtils)
	utils.Person{}.SayHello()
	//build类型为Package或者Directory,可以调用同一个包，但是不同文件下的函数及变量
	//build类型为File，只能调用当前文件及引入文件下的函数及变量
	res := add(1, 2)
	fmt.Println(res)
}
