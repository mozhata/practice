package user

import (
	"fmt"
	"practice/go/beedemo/models"
	"practice/go/beedemo/util"

	"github.com/astaxie/beego/orm"
)

// CreateUser insert user model to DB and return userID
func CreateUser(u models.User, orms ...orm.Ormer) (int64, error) {
	if !u.IsValid() {
		return 0, util.InvalidArgumentErr(nil, "user model is not valid")
	}
	var o orm.Ormer
	if len(orms) == 0 {
		o = orm.NewOrm()
	} else {
		o = orms[0]
	}
	uid, err := o.Insert(&u)
	if err != nil {
		return 0, util.InternalError(err, "insert user model to DB failed")
	}
	return uid, nil
}

func CheckExistance(userName string, orms ...orm.Ormer) (bool, error) {
	sql := fmt.Sprintf("select count(`id`) from %s where `name`=%q;", models.UserTable, userName)
	var count int
	err := models.QueryBySQL(sql, &count, false, orms...)
	if err != nil {
		return false, err
	}
	if count > 1 {
		return false, util.InternalError(nil, "count of user name %s should not more than 1 but it does. sql: %s", userName, sql)
	}
	return count == 1, nil
}

func AllUsers() ([]models.User, error) {
	var (
		all []models.User
		err error
	)
	sql := fmt.Sprintf("select * from %s;", models.UserTable)
	err = models.QueryBySQL(sql, &all, true)
	if err != nil {
		return nil, util.InternalError(err, "query by sql %s failed", sql)
	}
	return all, nil
}
