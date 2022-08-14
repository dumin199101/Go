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
  删除记录
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}
	//删除一条记录时，删除对象需要指定where条件,否则会全部删除
	// 1.先查询，再删除
	var user User
	db.Where("id", 15).First(&user)
	result := db.Delete(&user)
	fmt.Println("删除记录条数为：", result.RowsAffected)

	// 2.根据主键删除
	db.Delete(&User{}, 14)

	// 3.批量删除
	db.Delete(&User{}, []int{11, 12, 13})

	// 4.查找被软删除的记录(包括被软删除的)
	var users []User
	db.Unscoped().Where("name = @name", sql.Named("name", "jinzhu_2")).Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}
	var sql = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("name = @name", sql.Named("name", "jinzhu_2")).Find(&users)
	})
	fmt.Println(sql)

	// 5.永久删除
	db.Unscoped().Delete(&users)
}
