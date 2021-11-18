package main

import (
	"fmt"
	"io"
	"net/http"
)

// HTTP-客户端
func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Get Request Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println("Header:", resp.Header)
	fmt.Println("Proto:", resp.Proto)

	buffer := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buffer)
		if n == 0 {
			fmt.Println("数据读取完成")
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("数据读取出错...:", err)
			return
		}

		data := string(buffer[:n])
		fmt.Println(data)

		/*if err != nil && err == io.EOF{
			fmt.Println("数据读取到末尾...:",err)
			return
		}*/
	}
}
