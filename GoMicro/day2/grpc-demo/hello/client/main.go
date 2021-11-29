package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-demo/proto/hello"
)

func main() {
	conn, e := grpc.Dial("127.0.0.1:8989", grpc.WithInsecure())
	if e != nil {
		fmt.Println("Grpc Dial Error:", e)
		return
	}
	defer conn.Close()
	//初始化客户端
	client := hello.NewHelloClient(conn)
	//调用方法
	response, e := client.SayHello(context.Background(), &hello.HelloRequest{Name: "Gopher"})
	if e != nil {
		fmt.Println("Client SayHello Error:", e)
		return
	}

	fmt.Println("Info:", response.Message)

}
