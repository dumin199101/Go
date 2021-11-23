package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
  爬虫案例:
    https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/
*/
var wg sync.WaitGroup
var baseurl = "https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/"
var reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
var chanImageUrls = make(chan string, 150)
var chanTask = make(chan int, 5)

func main() {

	//获取网页内容,获取图片链接
	for i := 1; i < 6; i++ {
		wg.Add(1)
		go getHTMLPic(i)
	}

	// 3.任务统计协程，统计5个任务是否都完成，完成则关闭管道
	wg.Add(1)
	go CheckOK()

	//下载图片
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go downloadImgs()
	}

	wg.Wait()

}

// 任务统计协程，关闭管道
func CheckOK() {
	var count int
	for {
		<-chanTask
		count++
		if count == 5 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

func downloadImgs() {
	defer wg.Done()
	// 从管道里边取出共享数据
	for url := range chanImageUrls {
		//downloadPic(url)
		fmt.Println(url)
	}
}

//https://uploadfile.bizhizu.cn/up/d2/c6/ee/d2c6eea4c2b4c2fa030d8b506938d4b9.jpg
func downloadPic(url string) {
	n := strings.LastIndex(url, "/")
	name := url[n+1:]
	//使用时间戳解决重名问题
	timePrex := strconv.Itoa(int(time.Now().UnixNano()))
	name = timePrex + "_" + name
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Http Get Img Error:", err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	err = ioutil.WriteFile("D:\\irealer\\"+name, bytes, 0777)
	if err != nil {
		fmt.Println("Download Pic Error:", err)
		return
	} else {
		fmt.Println(name + "图片下载完成")
	}

}

func getHTMLPic(i int) {
	defer wg.Done()
	resp, err := http.Get(baseurl + strconv.Itoa(i) + ".html")
	if err != nil {
		fmt.Println("HTTP Get Error:", err)
		return
	}
	defer resp.Body.Close()
	buffer := make([]byte, 1024)
	result := ""
	for {
		n, err := resp.Body.Read(buffer)
		if n == 0 {
			fmt.Println("第" + strconv.Itoa(i) + "页数据下载完毕！")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("Read Error:", err)
			return
		}
		result += string(buffer[:n])
		//err = ioutil.WriteFile("D:\\irealer\\"+strconv.Itoa(i)+".txt", []byte(result), 0777)
		//if err != nil {
		//	fmt.Println("Write File Error:", err)
		//}
	}
	regex_img := regexp.MustCompile(reImg)
	results := regex_img.FindAllStringSubmatch(result, -1)
	for _, img := range results {
		url := img[0]
		chanImageUrls <- url
	}
	chanTask <- i
}
