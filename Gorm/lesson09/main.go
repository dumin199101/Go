package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
type User struct {
	Name string
	Age  int8
}

// 映射模型用
type UserInfo struct {
	gorm.Model
	Name     string
	Age      int8
	Birthday time.Time
}

// 存储分组结果
type Result struct {
	Age   string
	Total int
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}

func (UserInfo) TableName() string {
	return "user"
}

func main() {

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}

	// 7.选定特定的字段
	var users []User
	db.Select("name", "age").Find(&users)
	fmt.Println(users)

	// 8. order排序
	db.Where("name = ?", "hello").Order("age desc, name").Find(&users)
	fmt.Println(users)

	// 9.Limit & Offset
	// Cancel limit condition with -1
	var (
		users1 []User
		users2 []User
		users3 []User
	)
	db.Limit(2).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	fmt.Println(users1)
	fmt.Println(users2)

	db.Limit(3).Offset(5).Find(&users3)
	// SELECT * FROM users OFFSET 5 LIMIT 10;
	fmt.Println(users3)

	// 打印sql
	statement := db.Session(&gorm.Session{DryRun: true}).First(&UserInfo{}, map[string]interface{}{"name": "jinzhu"}).Statement
	println(statement.SQL.String())
	println(statement.Vars)
	println(db.Dialector.Explain(statement.SQL.String(), statement.Vars))

	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.First(&UserInfo{}, map[string]interface{}{"name": "jinzhu"})
	})
	println(sql)
}
