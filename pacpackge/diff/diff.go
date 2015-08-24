package main

import (
	"fmt"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	diffs()
}
func diffs() {
	// func Diff(A, B string) string
	// 行级比较,同"git diff"
	// 拿A跟B做比较,若A有,而B没有,则前置"-",若A没有而B有,则前置"+"
	constitution := strings.TrimSpace(`
We the People of the United States, in Order to form a more perfect Union,
establish Justice, insure domestic Tranquility, provide for the common defence,
promote the general Welfare, and secure the Blessings of Liberty to ourselves
and our Posterity, do ordain and establish this Constitution for the United
States of America.
`)

	got := strings.TrimSpace(`
:wq
We the People of the United States, in Order to form a more perfect Union,
establish Justice, insure domestic Tranquility, provide for the common defence,
and secure the Blessings of Liberty to ourselves
and our Posterity, do ordain and establish this Constitution for the United
States of America.
`)
	P(diff.Diff(got, constitution))
}
