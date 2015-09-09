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
type MyDB struct {
	DB *sql.DB
}

func NewDB(db *sql.DB) *MyDB {
	return &MyDB{DB: db}
}
func main() {
	// sql.Open(“mysql”, “user:password@/dbname”)
	db, e := sql.Open("mysql", "root:dx123!@/test")
	checkErr(e)
	defer db.Close()

	// 查看connection to database 是否仍然alive,若无,则重建链接
	e = db.Ping()
	checkErr(e)

	// Db := NewDB(db)
	// Db.insert("zhen", "zhenzhen", 333)

	// query by number
	smtOut, e := db.Prepare("select `id`, `username`, `password`, `number` from user where number=? order by id desc")
	checkErr(e)
	var u User
	e = smtOut.QueryRow(23).Scan(&u.Id, &u.UserName, &u.Pwd, &u.Number)
	checkErr(e)
	P(u)

	// db.Query
	rows, e := db.Query("select `username`, `password` from user order by id desc")
	columns, e := rows.Columns()
	checkErr(e)
	// defer rows.Close()
	P("columns: ", columns, len(columns))

	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(columns))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rows.Next() {
	// 	e = rows.Scan(scanArgs...)

	// 	for i, col := range values {
	// 		fmt.Println(columns[i], string(col))
	// 	}
	// 	fmt.Println("-----------------------------------")
	// }

}
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func (db *MyDB) insert(name, pwd string, num int) {
	smt, e := db.DB.Prepare("insert into user(`username`,`password`,`number`) values(?,?,?)")
	checkErr(e)
	defer smt.Close()

	result, e := smt.Exec(name, pwd, num)
	id, e := result.LastInsertId()
	checkErr(e)
	affected, e := result.RowsAffected()
	checkErr(e)
	P("LastInsertedId: ", id, "rowsAffected: ", affected)

}

// func (db *MyDB) ShowByNum(num int) {
// 	var u User
// 	smtOut, e := db.DB.Prepare("select")

// }
