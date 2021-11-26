package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

// 传的参数
type Params struct {
	Width, Height int
}

// rpc客户端
func main() {
	client, e := jsonrpc.Dial("tcp", "127.0.0.1:9999")
	if e != nil {
		fmt.Println("Dial HTTP Error:", e)
		return
	}
	zhouchang := 0
	e = client.Call("React.ZhouChang", Params{50, 20}, &zhouchang)
	if e != nil {
		fmt.Println("React.ZhouChang Error:", e)
		return
	}
	fmt.Println("矩形的周长是：", zhouchang)
	mianji := 0
	e = client.Call("React.MianJi", Params{50, 20}, &mianji)
	if e != nil {
		fmt.Println("React.MianJi Error:", e)
		return
	}
	fmt.Println("矩形的面积是：", mianji)
}
