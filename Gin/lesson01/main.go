package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type User struct {
	gorm.Model
	Name     string
	Age      int8
	Birthday time.Time
}

func (User) TableName() string {
	return "user"
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello,world",
		})
	})

	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		dsn := "root:root@tcp(127.0.0.1:3306)/fastadmin?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("数据库连接失败！%#v\n", err)
		}
		var users []Result
		db.Model(&User{}).Where("name like @name", sql.Named("name", name+"%")).Find(&users)
		if len(users) < 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": 404,
				"msg":  "对不起，没有找到合适的记录",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": users,
			"msg":  "数据查找成功",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("服务器启动失败，请稍后重试")
	}
}
