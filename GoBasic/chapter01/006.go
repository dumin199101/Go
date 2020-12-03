package main

import "fmt"

func main() {
	//键盘输入
	//一次接收一个
	var a int
	var name string
	fmt.Println("请输入年龄：")
	fmt.Scanln(&a)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("您输入的年龄是:%v,您输入的姓名是:%s", a, name)

	//一次接收多个
	fmt.Println("请输入年龄，姓名，用空格分隔")
	fmt.Scanf("%d %s", &a, &name)
	fmt.Printf("您输入的年龄是:%v,您输入的姓名是:%s", a, name)

}
