package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 模拟HTTP服务器，读取服务器文件

func main() {

	http.HandleFunc("/", cback)

	http.ListenAndServe("127.0.0.1:8090", nil)

}

func cback(writer http.ResponseWriter, request *http.Request) {
	basedir := "N:"
	filename := request.URL.Path
	file, e := os.Open(basedir + filename)
	if e != nil {
		fmt.Println("文件不存在！！！")
		return
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 2048)
	for {
		n, e := reader.Read(buffer)
		if n == 0 {
			fmt.Println("文件读取完成")
			return
		}
		if e != nil && e != io.EOF {
			fmt.Println("文件读取失败")
			return
		}
		writer.Write(buffer[:n])
	}
}
