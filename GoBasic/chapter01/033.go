package main

import (
	"fmt"
	"strconv"
)

type Lions struct {
	name    string
	age     int
	address string
}

func main() {
	//管道存取接口类型结构体，然后遍历
	lions1 := Lions{
		name:    "狮王",
		age:     1,
		address: "动物园",
	}

	lions2 := Lions{
		name:    "狮子王",
		age:     10,
		address: "森林",
	}

	channelLions := make(chan interface{}, 2)
	channelLions <- lions1
	channelLions <- lions2

	close(channelLions)
	for value := range channelLions {
		if v, ok := value.(Lions); ok {
			info := v.name + "," + strconv.Itoa(v.age) + "," + v.address
			fmt.Println(info)
		}
	}

}
