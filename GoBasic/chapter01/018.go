package main

import "fmt"

func main() {
	//切片
	var slice1 []int = []int{11, 22, 33}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 []string = make([]string, 4)
	slice2[0] = "wangyuanyuan"
	slice2[1] = "zhangzuolin"
	slice2[2] = "zhaoyazhi"
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2[0:2])
	slice2 = append(slice2, "zhangxueliang")
	for key, value := range slice2 {
		fmt.Println(key, value)
	}

	slice3 := []string{"hello", "world", "java"}
	for key, value := range slice3 {
		fmt.Println(key, value)
	}

}
