package main

import "fmt"

func main() {
	//数组的使用
	var arr0 [3]int
	fmt.Println(arr0)
	arr0[0] = 111
	arr0[1] = 222
	arr0[2] = 333
	fmt.Println(arr0)

	//arr0的地址:0xc000012340,arr0[0]的地址:0xc000012340,arr0[1]的地址:0xc000012348,arr0[2]的地址:0xc000012350
	fmt.Printf("arr0的地址:%p,arr0[0]的地址:%p,arr0[1]的地址:%p,arr0[2]的地址:%p\n", &arr0, &arr0[0], &arr0[1], &arr0[2])

	//数组初始化及遍历
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(len(arr))

	var arr2 = [3]int{4, 5, 6}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	var arr3 = [...]int{11, 12, 13, 14}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	var arr4 = [...]string{0: "lisi", 1: "wangwu", 2: "zhangsan"}
	for key, value := range arr4 {
		fmt.Println(key, value)
	}

}
