package main

import (
	"io"
	"net/http"
	"net/rpc"
)

/*
  golang写RPC程序，必须符合4个基本条件，不然RPC用不了

	结构体字段首字母要大写，可以别人调用

	函数名必须首字母大写

	函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型

	函数还必须有一个返回值error
*/

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
	http.HandleFunc("/panda", pandaFunc)
	//注册服务
	react := new(React)
	rpc.Register(react)
	//HTTP绑定
	rpc.HandleHTTP()
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func pandaFunc(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello,World!!!")
}
