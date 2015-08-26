package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var P func(...interface{}) (int, error) = fmt.Println

type User struct {
	Id       int
	UserName string
	Pwd      string
	Number   int
}

func main() {
	// func Open(driverName, dataSourceName string) (*DB, error)
	// 用户名:root,密码:dx123! 数据库:test
	db, e := sql.Open("mysql", "root:dx123!/test")
	checkErr(e)
	defer db.Close()


}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
