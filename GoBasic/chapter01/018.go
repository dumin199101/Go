package main

import "fmt"

func main() {
	//切片
	//切片是数组的一个引用，因此切片是引用类型。但自身是结构体，
	var slice1 []int = []int{11, 22, 33}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 []string = make([]string, 4)
	slice2[0] = "wangyuanyuan"
	slice2[1] = "zhangzuolin"
	slice2[2] = "zhaoyazhi"
	slice2[3] = "sunyang"
	//slice2[4] = "zhangwei" //切片越界
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2[0:2])
	slice2 = append(slice2, "zhangxueliang")
	for key, value := range slice2 {
		fmt.Println(key, value)
	}
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2)) // 容量扩充2倍

	slice3 := []string{"hello", "world", "java"}
	for key, value := range slice3 {
		fmt.Println(key, value)
	}

	fmt.Println("***************************")

	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := make([]int, 5)
	copy(slice5, slice4) //copy函数完成拷贝，数据空间独立，互不影响
	fmt.Println(slice4)
	fmt.Println(slice5)
	slice5[0] = 10
	fmt.Println(slice4)
	fmt.Println(slice5)

	fmt.Println("****************************")

	//切片的引用传递
	slice6 := []int{1, 2, 3, 4, 5}
	var slice7 = slice6
	slice7[0] = 10
	fmt.Println(slice6)
	fmt.Println(slice7)

}
