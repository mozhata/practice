package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// DB settings
	orm.RegisterDataBase("default", "mysql", "root:ddd@tcp(localhost:3306)/demo?charset=utf8", 30)

	orm.RegisterModel(NewUserModel())
}
