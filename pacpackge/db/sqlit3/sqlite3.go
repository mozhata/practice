package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func insert(db *sql.DB) (id int64) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("yongkang", "研发部门", "2012-12-09")
	checkErr(err)

	id, err = res.LastInsertId()
	checkErr(err)

	fmt.Println("insert--id: ", id)
	return id
}
func update(db *sql.DB, keyword string, id int64) {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec(keyword, id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("update--affect: ", affect)
}
func show(db *sql.DB) {
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid: ", uid, "username: ", username, "department: ", department,
			"created: ", created)
	}
}
func delete(db *sql.DB, id int64) {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("delete--affect: ", affect)
}
func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()
	insert(db)
	update(db, "peipei", 4)
	delete(db, 1)
	show(db)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
