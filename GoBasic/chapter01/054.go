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

	_, err = conn.Do("SET", "name", "lisi")
	if err != nil {
		fmt.Println("设置数据失败", err)
		return
	}

	//获取数据
	s, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("获取数据失败", err)
	}

	fmt.Println(s)

	//关闭Redis连接
	defer conn.Close()

}
