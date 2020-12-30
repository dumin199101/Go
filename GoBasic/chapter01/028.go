package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	//序列化与反序列化

	//1.结构体序列化
	bird := Bird{
		Name:  "燕子",
		Age:   1,
		Color: "黑色",
	}

	data, err := json.Marshal(bird)
	if err != nil {
		fmt.Println("error:json err")
	}
	fmt.Printf("转换完成的json字符串数据：%s\n", data)

	//2.结构体反序列化
	str := `{"name":"鹦鹉","age":11,"color":"白色"}`
	err = json.Unmarshal([]byte(str), &bird) //此处使用地址，地址传递，改变值
	if err != nil {
		fmt.Println("error:json parse err")
	}
	fmt.Println(bird)

	//3.map序列化
	bird2 := make(map[string]interface{})
	bird2["name"] = "鸽子"
	bird2["age"] = 2
	bird2["color"] = "白色"
	data2, err2 := json.Marshal(bird2)
	if err2 != nil {
		fmt.Println("error:json err")
	}
	fmt.Printf("转换完成的json字符串数据：%s\n", data2)

}
