package main

import (
	"fmt"
	"reflect"
)

func set_value(a interface{}) {
	value := reflect.ValueOf(a)
	kind := value.Kind()
	switch kind {
	case reflect.Float64:
		value.SetFloat(12.00) //reflect.Value.SetFloat using unaddressable value
		fmt.Println("a", value.Float())
	case reflect.Ptr:
		value.Elem().SetFloat(13.00)
		fmt.Println("a:", value.Elem().Float())
		fmt.Println("a的地址", value.Pointer())

	}
}

func main() {
	//反射修改值
	var x = 3.4
	set_value(&x)
	fmt.Printf("x的值变为：%.2f\n", x)
}
