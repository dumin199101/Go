package main

import "fmt"

/*交换变量*/
var mm int = 100
var nn int = 200

func main() {
	mm, nn = nn, mm
	fmt.Println(mm, nn)
}
