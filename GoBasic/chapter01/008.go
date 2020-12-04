package main

import "fmt"

func main() {
	//流程控制一：if else
	age := 22
	if age < 20 {
		fmt.Println("青少年时期")
	} else if age >= 20 && age < 40 {
		fmt.Println("青年时期")
	} else if age >= 40 && age < 60 {
		fmt.Println("中年时期")
	} else {
		fmt.Println("老年时期")
	}

	//流程控制二：switch,不用加break
	year := "2020"
	switch year {
	case "2020":
		fmt.Println("今年的年份是2020")
	case "2021":
		fmt.Println("今年的年份是2021")
	default:
		fmt.Println("今年的年份是待定")
	}

	//switch不加表达式
	switch {
	case age < 20:
		fmt.Println("青少年时期")
	case age >= 20 && age < 40:
		fmt.Println("青年时期")
	case age >= 40 && age < 60:
		fmt.Println("中年时期")
	default:
		fmt.Println("老年时期")
	}

	//switch穿透
	switch age1 := 18; {
	case age1 < 20:
		fmt.Println("青少年时期")
		fallthrough
	case age1 >= 20 && age1 < 40:
		fmt.Println("青年时期")
		fallthrough
	case age1 >= 40 && age1 < 60:
		fmt.Println("中年时期")
	default:
		fmt.Println("老年时期")
	}

}
