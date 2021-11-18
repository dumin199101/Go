package main

import (
	"fmt"
	"net/http"
)

// HTTP-服务端

func main() {
	http.HandleFunc("/hello", callback)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func callback(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Header:", request.Header)
	fmt.Println("Body:", request.Body)
	fmt.Println("Method:", request.Method)
	fmt.Println("RemoteAddr:", request.RemoteAddr)
	fmt.Println("Host:", request.Host)
	fmt.Println("URL:", request.URL)
	writer.Write([]byte("Hello,Welcom to My Website!!!"))
}
