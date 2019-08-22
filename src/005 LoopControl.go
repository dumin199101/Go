package main

import "fmt"

/**
  循环控制：
    1.go语言分支结构
    2.go语言循环结构
    3.go语言遍历
*/
var ten int = 11

func main() {
	// 分支结构
	if ten > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<=10")
	}

	//循环结构
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//输出乘法口诀表
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d", x, y, x*y)
		}
		fmt.Println()
	}

	//遍历数组
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d,vaule:%d", key, value)
	}

}
