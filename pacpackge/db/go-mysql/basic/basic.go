package main // package main

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
	// sql.Open(“mysql”, “user:password@/dbname”)
	db, e := sql.Open("mysql", "root:dx123!@/test")
	checkErr(e)
	defer db.Close()

	// 查看connection to database 是否仍然alive,若无,则重建链接
	e = db.Ping()
	checkErr(e)

	stmIns, e := db.Prepare("insert into user(`username`,`password`,`number`) values(?,?,?)")
	checkErr(e)

	result, e := stmIns.Exec("user1", "password1", 1)
	checkErr(e)

	m, e := result.LastInsertId()
	n, e := result.RowsAffected()
	P("last inser id: ", m, "rows affected: ", n)

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
