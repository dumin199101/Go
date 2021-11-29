package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-demo/proto/hello"
	"net"
)

type HelloService struct{}

func (h HelloService) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := new(hello.HelloResponse)
	response.Message = fmt.Sprintf("Hello %s.", in.Name)
	return response, nil
}

func main() {
	listener, e := net.Listen("tcp", "127.0.0.1:8989")
	if e != nil {
		fmt.Println("TCP Listen Error!!!!")
		return
	}

	helloService := HelloService{}

	//实例化grpc Server
	server := grpc.NewServer()
	// 注册服务
	hello.RegisterHelloServer(server, helloService)

	server.Serve(listener)
}
