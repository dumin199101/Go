package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //mysqlDriver时会用到, unknown driver "mysql" (forgotten import?)
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into user(username, sex, email)values(?, ?, ?)", "stu004", "man", "stu04@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)

	defer Db.Close() // 注意这行代码要写在上面err判断的下面
}
