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

var d Mysql

func main() {
	// sql.Open(“mysql”, “user:password@/dbname”)
	db, e := sql.Open("mysql", "root:dx123!@/test")
	checkErr(e)
	defer db.Close()

	// 查看connection to database 是否仍然alive,若无,则重建链接
	e = db.Ping()
	checkErr(e)
	// db.insert("name", "123", 123)
	d.DB = db
	d.insert("name", "pwd", 123)

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

type Mysql struct {
	DB *sql.DB
}

func (db *Mysql.DB) insert(name, pwd string, num int) {
	stmIns, e := db.Prepare("insert into user(`username`,`password`,`number`) values(?,?,?)")
	checkErr(e)
	defer stmIns.Close()

	result, e := stmIns.Exec(name, pwd, num)
	checkErr(e)

	m, e := result.LastInsertId()
	n, e := result.RowsAffected()
	P("last inser id: ", m, "rows affected: ", n)
}
