package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/mozhata/merr"
)

func QueryBySQL(sql string, container interface{}, multi bool, orms ...orm.Ormer) error {
	var o orm.Ormer
	if len(orms) == 0 {
		o = orm.NewOrm()
	} else {
		o = orms[0]
	}
	var err error
	if multi {
		_, err = o.Raw(sql).QueryRows(container)
	} else {
		err = o.Raw(sql).QueryRow(container)
	}
	if err != nil {
		return merr.InternalError(err, "query by sql %s failed", sql)
	}
	return nil
}

func ExcueSQL(sql string, orms ...orm.Ormer) error {
	var o orm.Ormer
	if len(orms) != 1 {
		o = orm.NewOrm()
	} else {
		o = orms[0]
	}
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return merr.InternalError(err, "excute sql %s failed", sql)
	}
	return nil
}
