package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 参数
type Params struct {
	Width  int
	Height int
}

// 矩形
type React struct{}

func (this *React) ZhouChang(params Params, ret *int) error {
	*ret = (params.Width + params.Height) * 2
	return nil
}

func (this *React) MianJi(params Params, ret *int) error {
	*ret = params.Width * params.Height
	return nil
}

func main() {
	//注册服务
	react := new(React)
	rpc.Register(react)
	listener, e := net.Listen("tcp", "127.0.0.1:9999")
	if e != nil {
		fmt.Println("Net Listen Error:", e)
		return
	}
	for {
		// 监听客户端连接
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("Listener Accept Error:", e)
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
