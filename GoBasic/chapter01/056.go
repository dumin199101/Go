package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	//Redis 连接池
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		MaxIdle:     8,   //Redis连接池默认空闲时连接数
		MaxActive:   0,   //Redis连接池最大连接数
		IdleTimeout: 300, //连接不使用300s关闭
	}
}

func main() {
	//获取连接
	conn := pool.Get()
	//使用完毕，释放连接
	defer conn.Close()

	_, err := conn.Do("HMSET", "user001", "name", "wangwei", "age", 22)
	if err != nil {
		fmt.Println("设置数据失败")
	}

	strings, err := redis.Strings(conn.Do("HMGET", "user001", "name", "age"))
	if err != nil {
		fmt.Println("获取数据失败")
	}

	for key, value := range strings {
		fmt.Println(key, value)
	}

	//使用完毕，关闭连接池
	pool.Close()
}
