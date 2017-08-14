package user

import (
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
