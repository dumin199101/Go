package main

import (
	"fmt"
	"reflect"
)

type Worker struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (w Worker) Sum(a int, b int) int {
	return a + b
}

func test_struct(a interface{}) {
	//获取类型
	t := reflect.TypeOf(a)
	fmt.Println(t)
	fmt.Println(t.Name())
	//获取值
	value := reflect.ValueOf(a)
	fmt.Println(value)
	//将value转为interface{}
	ivalue := value.Interface()
	//类型断言
	worker, ok := ivalue.(Worker)
	if ok {
		fmt.Println(worker.Age, worker.Name, worker.Address)
	}

	//获取所有属性
	for i := 0; i < t.NumField(); i++ {
		//t.Field(i) --- StructField
		fmt.Printf("字段的名称%s,字段的类型%v,标签的值是：%s\n", t.Field(i).Name, t.Field(i).Type, t.Field(i).Tag)
		//reflect.Value类型转换interface{},此时应注意属性的访问权限：cannot return value obtained from unexported field or method
		ivalue := value.Field(i).Interface()
		if v, ok := ivalue.(string); ok {
			fmt.Println("[string]val :", v)
		}

		if v, ok := ivalue.(int); ok {
			fmt.Println("[int]val :", v)
		}

		//标签值
		s := t.Field(i).Tag.Get("json")
		fmt.Printf("标签%s的值为%s\n", t.Field(i).Tag, s)
	}

	fmt.Println("方法：")
	//获取方法
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name)
		fmt.Println(method.Type)
	}

	//调用方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(30))
	res := value.Method(0).Call(params)
	fmt.Println(res[0].Int()) //返回值[]reflect.Value

}

func main() {
	//反射与结构体
	test_struct(Worker{Name: "lisi", Age: 22, Address: "jinan"})

}
