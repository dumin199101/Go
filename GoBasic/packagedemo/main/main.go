package main

import (
	"Go/GoBasic/packagedemo/utils"
	"fmt"
)

func main() {
	// 引入其他包的变量，默认从GoPath的src目录开始寻址，调用方式：包名.变量名
	fmt.Print(utils.DbUtils)
}
