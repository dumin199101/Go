package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 切片指定len跟cap：len=high-low cap=max-low
	var arr1 = arr[1:5:9]  // low:high:max
	fmt.Println(arr1)      // 2,3,4,5
	fmt.Println(len(arr1)) // 5-1 4
	fmt.Println(cap(arr1)) // 9-1 8

	//切片未指定cap cap的长度就是数组的长度或者切片的cap-low
	var arr2 = arr[0:5]
	fmt.Println(arr2)      // 1,2,3,4,5
	fmt.Println(len(arr2)) //5
	fmt.Println(cap(arr2)) // 10

	var arr3 = arr1[1:3]   // 2,3,4,5
	fmt.Println(arr3)      //3,4
	fmt.Println(len(arr3)) // 3-1 2
	fmt.Println(cap(arr3)) // 8-1 7
}
