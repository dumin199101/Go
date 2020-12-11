package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串长度
	str := "hello北京"
	fmt.Println("字符串的长度:", len(str))
	//字符串遍历
	str = "hello,中国"
	for key, value := range str {
		fmt.Printf("key:%v,value:%c\n", key, value)
	}
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("value:%c\n", str1[i])
	}
	//字符串整数相互转换
	str2 := "12"
	i1, _ := strconv.Atoi(str2)
	fmt.Println(i1)
	i2 := 100
	str3 := strconv.Itoa(i2)
	fmt.Println(str3)
	//字符串与byte切片互转
	str4 := "hello"
	var bytes = []byte(str4)
	fmt.Printf("value:%v\n", bytes)
	var str5 = string([]byte{97, 98, 99, 100})
	fmt.Printf("value:%s\n", str5)
	//十进制与其他进制转换
	var str6 = strconv.FormatInt(10, 2)
	fmt.Println(str6)
	//查找指定子串
	var flag = strings.Contains("hello", "ll")
	fmt.Println(flag)
	//统计子串出现次数
	var count = strings.Count("hello", "l")
	fmt.Println(count)
	//子串第一次出现位置
	var index = strings.Index("hello", "e")
	fmt.Println(index)
	var index2 = strings.LastIndex("hello", "l")
	fmt.Println(index2)
	//替换
	var str7 = strings.Replace("hello world,hello beijing", "hello", "welcome", -1)
	fmt.Println(str7)
	//分割
	strArr := strings.Split("hello,world,go", ",")
	for key, value := range strArr {
		fmt.Println(key, value)
	}
	//大小写转换
	str8 := strings.ToLower("ABC")
	fmt.Println(str8)
	str9 := strings.ToUpper("abc")
	fmt.Println(str9)
	//判断前缀后缀
	flag2 := strings.HasPrefix("http://www.baidu.com", "http")
	fmt.Println(flag2)
	flag3 := strings.HasSuffix("10.jpg", "jpg")
	fmt.Println(flag3)
	//去掉空白或指定字符
	str10 := strings.TrimSpace(" hello,world  ")
	fmt.Println(str10)
	str11 := strings.Trim("hello,world!!!!", "!")
	fmt.Println(str11)

}
