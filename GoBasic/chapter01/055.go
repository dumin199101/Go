package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Redis Connection err", err)
		return
	}

	fmt.Println("Redis Connection success")

	//进行批量设置

	_, err = conn.Do("MSET", "age", 22, "salary", 10000)
	if err != nil {
		fmt.Println("设置数据失败", err)
		return
	}

	//获取数据
	data, err := redis.Ints(conn.Do("MGET", "age", "salary"))
	if err != nil {
		fmt.Println("获取数据失败", err)
	}

	for _, value := range data {
		fmt.Println(value)
	}

	//关闭Redis连接
	defer conn.Close()

}
