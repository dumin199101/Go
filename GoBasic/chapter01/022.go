package main

import "fmt"

type Personal struct {
	Name string
}

func (p Personal) study() {
	fmt.Println("好好学习")
}

func (p *Personal) study2() {
	fmt.Println("天天向上")
}

func main() {
	//对于函数，接收者为值类型，不能使用指针类型直接传递给方法，反之亦然
	//对于方法，接收者为值类型，可以使用指针类型变量调用方法，反之亦然
	person := Personal{Name: "张三"}
	person.study()
	person.study2()
	person2 := &Personal{Name: "李四"}
	person2.study()
	person2.study2()
}
