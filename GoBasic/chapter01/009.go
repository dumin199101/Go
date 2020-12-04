package main

import "fmt"

func main() {
	//循环语句,Go语言中没有while循环跟do-while循环
	// 1.for循环
	for i := 1; i < 10; i++ {
		fmt.Println("Hello World")
	}

	// 2.for-range循环
	str := "helloworld"

	for key, value := range str {
		fmt.Printf("key:%d,value:%c\n", key, value)
	}

	// 3.break语句
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i:%d", i)
	}

	fmt.Println("\n******************")

	//4.continue
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("i:%d", i)
	}

	//goto语句
	fmt.Println("开始游戏")
	num := 5
	fmt.Println("1....")
	fmt.Println("2....")
	fmt.Println("3....")
	fmt.Println("4....")
	if num == 5 {
		goto label3
	}
	fmt.Println("5....")
	fmt.Println("6....")
label3:
	fmt.Println("7....")
	fmt.Println("8....")
	fmt.Println("9....")
	fmt.Println("10....")

}
