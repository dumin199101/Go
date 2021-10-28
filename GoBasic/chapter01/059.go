package main

import "fmt"

// 修改字符串

func main() {
	s1 := "hello"
	arr1 := []byte(s1)
	arr1[0] = 'w'
	s1 = string(arr1)
	fmt.Println(s1)

	s2 := "北京欢迎你"
	arr2 := []rune(s2)
	arr2[0] = '南'
	s2 = string(arr2)
	fmt.Println(s2)
}
