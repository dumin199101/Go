package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
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
  查询
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}

	// 1.查询单条记录
	var user User
	db.Where("name = @name", sql.Named("name", "王芳")).First(&user)
	fmt.Println(user)

	// 2.获取最后一条符合条件数据
	var user2 User
	db.Last(&user2)
	fmt.Println(user2)

	// 3.使用主键检索
	var user3 User
	db.First(&user3, 8)
	fmt.Println(user3)

	// 4.检索全部记录
	var users []User
	db.Find(&users, []int{1, 3, 6, 8, 9, 10})
	for i, user := range users {
		fmt.Println(i, user.ID, user.Name)
	}

	// 5.where-string条件

	// Get first matched record
	var user4 User
	db.Where("name = ?", "hello").First(&user4)
	fmt.Println("user4:", user4)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// Get all matched records
	var users5 []User
	db.Where("name <> ?", "jinzhu").Find(&users5)
	fmt.Println("users5:", users5)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	var users6 []User
	db.Where("name IN ?", []string{"hello", "jinzhu_2"}).Find(&users6)
	fmt.Println("users6", users6)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	db.Where("updated_at > ?", time.Now().Add(-1*time.Hour)).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("age BETWEEN ? AND ?", 20, 25).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// 6. where - struct and map and slice条件 [注意：当使用struct作为查询条件时会忽略零值，如果想使用零值参与查询，使用map]

	// Struct
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Slice of primary keys
	db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

}
