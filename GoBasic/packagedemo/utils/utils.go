package utils

import "fmt"

var DbUtils string = "DbUtils"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Println("Hello")
}
