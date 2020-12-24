package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Student struct {
	name     string
	age      int8
	address  string
	isFemale bool
	school   map[string]string
}

type Person struct {
	name string
	age  int8
}

func newPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) dream() {
	fmt.Printf("%s的梦想是:学好Go语言。\n", p.name)
}

type Animal struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func (p Person) setAge(age int8) {
	p.age = age
}

func (p *Person) setAge2(age int8) {
	p.age = age
}

func (p *Person) String() string {
	str := "[" + p.name + "," + strconv.Itoa(int(p.age)) + "]"
	return str
}

// 结构体匿名字段
type Dog struct {
	string
	int8
}

//结构体嵌套
type Player struct {
	name string
	age  int8
}

type BasketballPlayer struct {
	gender  string
	address string
	Player
}

func main() {
	//struct 结构体
	//初始化

	//方式一
	student := Student{
		name:     "张三",
		age:      20,
		address:  "北京市",
		isFemale: false,
		school: map[string]string{
			"junior":  "中学生",
			"senior":  "高中生",
			"college": "大学生",
		},
	}

	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.address)
	fmt.Println(student.isFemale)
	for key, value := range student.school {
		fmt.Println(key, value)
	}

	//方式二
	var student2 Student
	student2.name = "李四"
	student2.age = 22
	fmt.Println(student2.name)
	fmt.Println(student2.age)

	//方式三
	var student3 *Student = new(Student)
	student3.name = "王五"
	student3.age = 39
	fmt.Println(student3.name)
	fmt.Println(student3.age)

	//方式四
	var student4 *Student = &Student{
		name:     "朱六",
		age:      30,
		address:  "天津市",
		isFemale: false,
		school: map[string]string{
			"junior": "中学生",
			"senior": "中专生",
		},
	}

	fmt.Println(student4.name)
	fmt.Println(student4.age)
	fmt.Println(student4.address)
	fmt.Println(student4.isFemale)
	for key, value := range student4.school {
		fmt.Println(key, value)
	}

	//匿名结构体
	var student5 = struct {
		name string
		age  int8
	}{
		name: "匿名者1",
		age:  33,
	}

	fmt.Println(student5.name)
	fmt.Println(student5.age)

	//内存布局:内存地址是连续的
	var test = struct {
		a int8
		b int8
		c int8
		d int8
	}{
		a: 1,
		b: 2,
		c: 3,
		d: 4,
	}

	fmt.Printf("test.a的内存地址：%p\n", &test.a)
	fmt.Printf("test.b的内存地址：%p\n", &test.b)
	fmt.Printf("test.c的内存地址：%p\n", &test.c)
	fmt.Printf("test.d的内存地址：%p\n", &test.d)

	//工厂模式实例化对象
	person := newPerson("郑强", 19)
	fmt.Println(person.name)
	fmt.Println(person.age)

	//结构体方法
	person.dream()

	//结构体标签
	animal := Animal{
		Name: "Cat",
		Age:  1,
	}
	data, e := json.Marshal(animal)
	if e != nil {
		fmt.Println("error:", e)
		return
	}
	fmt.Printf("返回的json数据：%s\n", data)

	//方法参数的值传递跟引用传递的比较
	person2 := Person{
		name: "张三",
		age:  0,
	}

	person2.setAge(12)
	fmt.Println(person2.age) //0
	person2.setAge2(20)
	fmt.Println(person2.age) //20

	//结构体重写String方法,fmt.Println会调用String方法返回值
	fmt.Println(&person2)

	//结构体匿名字段
	dog := Dog{
		"旺财",
		2,
	}

	fmt.Printf("%v的年龄是%d岁\n", dog.string, dog.int8)

	//结构体嵌套,类似于继承
	basketballPlayer := BasketballPlayer{
		gender:  "男",
		address: "美国洛杉矶",
		Player: Player{
			name: "james",
			age:  22,
		},
	}

	fmt.Printf("篮球运动员的姓名是%v,年龄是%d,性别是%v,地址是%v\n", basketballPlayer.name, basketballPlayer.age, basketballPlayer.gender, basketballPlayer.address)

}
