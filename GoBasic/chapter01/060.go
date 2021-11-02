package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func runner(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("任务%v结束退出\n", ctx.Value(key))
			return
		default:
			fmt.Printf("任务%v正在运行中\n", ctx.Value(key))
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	//管理启动的协程
	ctx, cancel := context.WithCancel(context.Background())
	// 给ctx绑定键值，传递给goroutine
	valuectx := context.WithValue(ctx, key, "【监控1】")

	// 开启goroutine，传入ctx
	go runner(valuectx)

	// 运行一段时间后停止
	time.Sleep(time.Second * 10)
	fmt.Println("停止任务")
	cancel() // 使用context的cancel函数停止goroutine

	// 为了检测监控过是否停止，如果没有监控输出，表示停止
	time.Sleep(time.Second * 3)
}
