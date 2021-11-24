package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	prototest "testproto/prototest"
)

func main() {
	person := &prototest.Person{
		Name:   "Amy",
		Age:    22,
		Height: []int32{177},
		Motto:  "Hello Protobuf!",
	}

	fmt.Println(person)
	//编码
	bytes, e := proto.Marshal(person)

	if e != nil {
		fmt.Println("Marshal Error:", e)
	}

	fmt.Println(bytes) //二进制数据

	person1 := &prototest.Person{}
	//解码
	e = proto.Unmarshal(bytes, person1)
	if e != nil {
		fmt.Println("UnMarshal Error:", e)
	}
	fmt.Println(person1.GetName())
	fmt.Println(person1.GetAge())
	fmt.Println(person1.GetHeight())
	fmt.Println(person1.GetMotto())

}
