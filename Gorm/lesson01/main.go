package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Userinfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	conn, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("数据库连接失败！")
	}
	defer conn.Close()

	// 自动迁移
	conn.AutoMigrate(&Userinfo{})

	u1 := Userinfo{Id: 1, Name: "张三", Gender: "男", Hobby: "篮球"}
	u2 := Userinfo{Id: 2, Name: "李四", Gender: "女", Hobby: "篮球"}
	u3 := Userinfo{Id: 3, Name: "王五", Gender: "男", Hobby: "足球"}

	// 添加
	conn.Create(u1)
	conn.Create(u2)
	conn.Create(u3)

	// 查询
	user := new(Userinfo)
	conn.First(user)
	fmt.Printf("第一条记录为%#v\n", user)

	user2 := new(Userinfo)
	conn.Find(user2, "hobby=?", "足球")
	fmt.Printf("兴趣爱好是足球的为：%#v\n", user2)

	// 修改
	conn.Model(&user).Update("hobby", "羽毛球")

	// 删除
	conn.Delete(&user)

}
