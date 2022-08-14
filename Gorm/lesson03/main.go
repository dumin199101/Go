package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

/**
  数据库连接，简单的增删改查
*/
func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败！%#v\n", err)
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println("查询product的记录：:", product)
	creat_at := product.CreatedAt.Format("2006-01-02 15:03:04")
	fmt.Println("创建时间为：", creat_at)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)

}
