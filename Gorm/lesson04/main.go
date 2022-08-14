package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int8
	Birthday time.Time `gorm:"autoCreateTime"`
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}

/**
  创建记录
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}

	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建

	fmt.Println("插入数据库的主键：", user.ID)
	fmt.Println("返回的error：", result.Error)
	fmt.Println("返回插入记录的条数：", result.RowsAffected)

	// 批量插入:结构体切片方式
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	result = db.Create(&users)
	fmt.Println("结构体切片方式批量插入的记录数为：", result.RowsAffected)
	fmt.Println("返回的error：", result.Error)

	for _, user := range users {
		fmt.Println("插入的ID为：", user.ID)
	}

	// 批量插入：Map切片方式
	result = db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
	fmt.Println("Map切片方式批量插入的记录数为：", result.RowsAffected)

}
