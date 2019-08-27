package main

import (
	"fmt"
	"time"
)

/**
  复合数据类型
      1.数组
      2.散列表
      3.结构体
      4.切片
*/

// 1.数组声明
var gg [3]int

//数组字面量形式初始化数组
var gg2 = [...]int{1, 2, 3, 4, 5, 6}

//切片
var gg3 = []int{22, 33, 44, 55, 66, 77, 88}

//结构体
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	//数组
	fmt.Println(gg[0])
	fmt.Println(len(gg))
	for i, v := range gg {
		fmt.Printf("%d %d \n", i, v)
	}
	for _, v := range gg {
		fmt.Printf("%d\n", v)
	}
	fmt.Printf("%T\n", gg2)
	//散列表
	//1.使用make函数创建map
	ages := make(map[string]int)
	ages["hello"] = 22
	ages["world"] = 23
	ages["go"] = 24
	for key, value := range ages {
		fmt.Println(key, value)
	}

	//2.使用map字面量进行初始化
	names := map[string]int{
		"hello": 22,
		"world": 23,
		"go":    24,
	}
	for key, value := range names {
		fmt.Println(key, value)
	}
	//删除散列表中的元素
	delete(ages, "go")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//结构体
	dilbert.ID = 1
	dilbert.Name = "John"
	dilbert.Address = "北京市"
	fmt.Println(dilbert.Name)

	//获取成员变量的地址，通过指针操作
	salary := &dilbert.Salary
	*salary = *salary + 1000
	println(dilbert.Salary)

	//切片
	for k, v := range gg3 {
		fmt.Println(k, v)
	}

}
