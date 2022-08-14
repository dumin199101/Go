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
  修改数据
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}

	// 查询数据
	var user User
	db.First(&user, 4)
	location, err := time.LoadLocation("Asia/Shanghai")
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "1999/08/04 14:15:20", location)
	if err != nil {
		log.Fatalf("时区解析错误：%#v", err)
	}
	// 更新多个列(结构体解析)
	res := db.Model(&user).Updates(User{Name: "王芳", Birthday: timeObj})
	fmt.Println("更新的记录数为：", res.RowsAffected)

	// 更新多个列（map）[注意：被软删除的记录是不会被更新的]
	res = db.Model(&User{}).Where("id", 5).Updates(map[string]interface{}{"name": "赵匡", "age": 22})
	fmt.Println("更新的记录数为：", res.RowsAffected)

	// 批量更新
	res = db.Table("user").Where("id IN ?", []int{8, 9}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	fmt.Println("更新的记录数为：", res.RowsAffected)
}
