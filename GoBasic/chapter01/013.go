package main

import (
	"fmt"
	"time"
)

func main() {
	//时间日期
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("时间类型：%T\n", now)
	//获取时间
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//时间常量
	fmt.Println(time.Millisecond)
	fmt.Println(time.Second)
}
