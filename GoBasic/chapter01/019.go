package main

import "fmt"

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

}
