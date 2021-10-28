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
	fmt.Println(kind) // int,struct
	value := reflect.ValueOf(b)
	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("reflect.Struct的值：", reflect.Struct) //struct
		vi := value.Interface()
		if v, ok := vi.(Cc); ok {
			fmt.Printf("姓名：%s,年龄：%d", v.name, v.age)
		}
	case reflect.Int:
		fmt.Println("reflect.Int的值：", reflect.Int) //int
		fmt.Println("值", value)
		i := value.Int()
		fmt.Println(i)
	}

}

func main() {
	//反射:程序运行时对程序的访问和修改能力

	// 反射获取空接口类型信息及值信息
	reflect_test(10) //int
	fmt.Println("********************************")
	reflect_test(Cc{name: "cat", age: 22}) //struct

}
