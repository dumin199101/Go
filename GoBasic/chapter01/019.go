package main

import (
	"fmt"
	"sort"
)

func main() {
	//map
	var map1 map[string]int = make(map[string]int, 3)
	map1["age"] = 22
	map1["sex"] = 0
	map1["salary"] = 1000

	for key, value := range map1 {
		fmt.Println(key, value)
	}

	map2 := map[string]string{
		"hello": "gaga",
		"world": "haha",
		"www":   "guagua",
	}

	for key, value := range map2 {
		fmt.Println(key, value)
	}

	//删除元素
	delete(map2, "www")
	for key, value := range map2 {
		fmt.Println(key, value)
	}

	//map的有序输出（面试题）
	map3 := map[int]string{
		1: "关羽",
		4: "马超",
		3: "赵云",
		5: "黄忠",
		2: "张飞",
	}

	slice1 := []int{}

	for key, _ := range map3 {
		slice1 = append(slice1, key)
	}

	sort.Ints(slice1)

	for _, value := range slice1 {
		fmt.Println(value, map3[value])
	}

}
