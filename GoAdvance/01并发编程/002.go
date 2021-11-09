package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("GGGGGGGGGGGGGG")
		fmt.Println("AAAAAAAAAAAAAAA")
		test()
		defer fmt.Println("BBBBBBBBBBBBBBBBBB")
	}()

	time.Sleep(time.Second)
}

func test() {
	defer fmt.Println("EEEEEEEEEEEEEEEEEEE")
	fmt.Println("CCCCCCCCCCCCCCCC")
	defer fmt.Println("DDDDDDDDDDDDDDDDD")

}
