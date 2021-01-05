package main

import (
	"fmt"
	"reflect"
)

type Cc struct {
	name string
	age  int
}

func reflect_test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println("类型", t) //int,main.Cc
	kind := t.Kind()
	fmt.Println(kind)                          // int,struct
	fmt.Println("reflect.Int的值：", reflect.Int) //int
	value := reflect.ValueOf(b)
	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("值", value)
	case reflect.Int:
		fmt.Println("值", value)
		i := value.Int()
		fmt.Println(i)
	}

}

func main() {
	//反射:程序运行时对程序的访问和修改能力

	// 反射获取空接口类型信息及值信息
	reflect_test(10)                       //int
	reflect_test(Cc{name: "cat", age: 22}) //struct

}
