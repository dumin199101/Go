package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db1 *sqlx.DB

type User struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func init() {
	database, e := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if e != nil {
		fmt.Println("open mysql failed,", e)
		return
	}
	Db1 = database
}

func main() {
	var user []User
	err := Db1.Select(&user, "select user_id, username, sex, email from user where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
	}

	fmt.Println("select succ:", user)

	defer Db1.Close()

}
