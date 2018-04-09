package account

import (
	"practice/go/encrypt/server/util/encrypt"
	"unicode"

	"practice/go/encrypt/merr"
)

func password(uuid, pwd string) string {
	salt := "openTheGate"
	return encrypt.MD5Sum(uuid[8:], pwd, salt)[:8]
}

func CheckPasswordStrength(strs string) error {
	var (
		total    int
		numberic int
		upper    int
		lower    int
		graphic  int
	)

	for _, v := range strs {
		if unicode.IsUpper(v) {
			upper++
		} else if unicode.IsLower(v) {
			lower++
		} else if unicode.IsNumber(v) {
			numberic++
		} else if unicode.IsGraphic(v) {
			graphic++
		}
		total++
	}
	if total < 6 {
		return merr.WrapErr(nil, "密码太短")
	}
	if total > 18 {
		return merr.WrapErr(nil, "密码太长")
	}
	if upper < 1 {
		return merr.WrapErr(nil, "必须有大写字符")
	}
	if lower < 1 {
		return merr.WrapErr(nil, "必须有小写字符")
	}
	if numberic < 1 {
		return merr.WrapErr(nil, "必须有数字")
	}
	return nil
}
