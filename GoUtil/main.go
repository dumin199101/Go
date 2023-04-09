package main

import (
	"fmt"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/timex"
	"time"
)

func main() {
	toInt, _ := goutil.ToInt("120")
	fmt.Println(toInt)
	ret, _ := arrutil.ToInt64s([]string{"11", "22"})
	fmt.Println(ret)
	now := timex.Now()
	fromTime := timex.FromTime(time.Now())
	format := now.DateFormat("Y-m-d H:i:s")
	fmt.Println(now, fromTime, format)

}
