package models

import (
	"practice/go/beedemo/util"

	"github.com/astaxie/beego/orm"
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
		return util.InternalError(err, "query by sql %s failed", sql)
	}
	return nil
}
