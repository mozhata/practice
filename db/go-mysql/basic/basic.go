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
	db, e := sql.Open("mysql", "root:dx123!/test")
	if e != nil {
		panic(e)
	}
	defer db.Close()

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
