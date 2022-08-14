package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 存储结果
type Result struct {
	ID   int
	Name string
	Age  int
}

// 数据模型
type User struct {
	ID   int
	Name string
	Age  int
}

// 将 User 的表名设置为 `profile`
func (User) TableName() string {
	return "user"
}

/**
  原生SQL查询:Raw,把结果通过Scan方法扫描进变量
  原生SQL插入、修改、删除：Exec
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}
	// 结构体存储返回的结果
	var result []Result
	db.Raw("SELECT id, name, age FROM user WHERE name = ?", "jinzhu_2").Find(&result)
	fmt.Println("返回的记录为：", result)
	for index, res := range result {
		fmt.Println("数据索引：", index, ",ID：", res.ID, ",Age", res.Age, ",Name:", res.Name)
	}

	// 普通变量存储返回的结果
	var age int
	db.Raw("SELECT SUM(age) FROM user WHERE id in ?", []int{1, 2, 3}).Scan(&age)
	fmt.Println("sum(age):", age)

	// 使用命名参数：sql.Named形式
	var age2 int
	db.Raw("SELECT SUM(age) FROM user WHERE id in @ids", sql.Named("ids", []int{1, 2, 3, 4})).Scan(&age2)
	fmt.Println("sum(age):", age2)

	// 使用命名参数：map形式
	var age3 int
	db.Raw("SELECT SUM(age) FROM user WHERE name = @name1 OR name = @name2", map[string]interface{}{"name1": "Jinzhu", "name2": "jinzhu_2"}).Scan(&age3)
	fmt.Println("sum(age):", age3)

	// toSQL调试
	var sql = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id = ?", 1).Limit(10).Order("age desc").Find(&[]Result{})
	})
	fmt.Println(sql)

	// exec
	db.Exec("UPDATE user SET deleted_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

}
