package main

import "fmt"

/**
  循环控制：
    1.go语言分支结构
    2.go语言循环结构
    3.go语言range遍历
    4.go语言switch语句:不需要使用break语句跳出
    5.go语言goto语句
    6.go语言break语句
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

	//类似while循环
	var i = 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//while无限循环
	var n = 1
	for {
		if n > 3 {
			break
		}
		fmt.Println("无限循环", n)
		n++
	}

	//遍历数组
	for key, value := range [5]int{100, 200, 300, 400, 500} {
		fmt.Printf("key:%d,value:%d", key, value)
		fmt.Println()
	}

	//遍历切片
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d,vaule:%d", key, value)
		fmt.Println()
	}

	//遍历map
	m := map[string]int{
		"hello": 100,
		"world": 200,
		"go":    300,
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	// switch 语句
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println("Hello")
	case "world":
		fmt.Println("World")
	default:
		fmt.Println("Else")
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				fmt.Println(y)
				goto breakTag
			}
		}
	}

breakTag:
	fmt.Println("Go to Outside")

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}

	fmt.Println("------------------------")

OuterLoop2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 1:
				fmt.Println(i, j)
				continue OuterLoop2
			}
		}
	}

	fmt.Println("--------------------")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i > 5 {
			break
		}
		println(i)
	}

}
