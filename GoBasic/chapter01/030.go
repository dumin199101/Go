package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//文件操作

	//打开一个只读文件
	file, e := os.Open("D:\\google\\go\\1.txt")
	if e != nil {
		fmt.Println("文件打开失败")
	}
	fmt.Printf("%v\n", file)
	//读取文件内容
	reader := bufio.NewReader(file)
	for {
		str, e := reader.ReadString('\n') //读到换行就结束
		fmt.Print(str)
		if e == io.EOF { //文件末尾
			break
		}
	}
	e = file.Close()
	if e != nil {
		fmt.Println("文件关闭失败")
	}

	//文件不大的情况下一次性读取
	data, e := ioutil.ReadFile("D:\\google\\go\\1.txt")
	if e != nil {
		fmt.Println("文件读取失败")
	}
	fmt.Printf("\n%s", data)

	//写入文件
	openFile, e := os.OpenFile("D:\\google\\go\\2.txt", os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		fmt.Println("文件创建失败")
	}
	writer := bufio.NewWriter(openFile)
	writer.WriteString("Hello,你好\n")
	writer.WriteString("Go go\n")
	writer.Flush()
	e = openFile.Close()
	if e != nil {
		fmt.Println("文件关闭失败")
	}
	//直接向文件中写入
	e = ioutil.WriteFile("D:\\google\\go\\2.txt", []byte("Hello,World"), 0666) //清空写
	if e != nil {
		fmt.Println("写入失败")
	}

}
