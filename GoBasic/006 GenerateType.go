package main

import "fmt"

/**
  派生类型：
      1.数组
      2.切片
      3.Map
*/

func main() {
	//数组
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	//数组初始化
	var b [3]int = [3]int{1, 2, 3}
	for _, v := range b {
		fmt.Println(v)
	}

	//数组长度根据初始化值来计算
	c := [...]int{1, 2, 3, 4, 5}
	for k, v := range c {
		fmt.Println(k, v)
	}

	//多维数组
	array := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	//访问数组元素
	fmt.Println(array[1][1])

	//2.切片
	s := []int{50, 30, 20, 10, 43}
	fmt.Println(s[1:3])

	//使用append为切片添加元素
	s = append(s, 88, 99)
	fmt.Println(s)

	//Map映射
	map1 := map[string]int{
		"hello": 1,
		"world": 2,
		"go":    3,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//删除map元素
	delete(map1, "go")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

}
