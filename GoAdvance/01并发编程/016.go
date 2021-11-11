package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //随机数种子，每次结果都不同
	i := rand.Int()                  //随机数
	fmt.Println(i)
	intn := rand.Intn(100) //返回一个0-100之间的随机数
	fmt.Println(intn)
}
