package main

import (
	"fmt"

	"github.com/kylelemons/godebug/diff"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	diffs()
}
func diffs() {
	// func Diff(old, new string) string
	// 行级比较,同"git diff"
	// new 跟 old 比较,
	// -old
	// +new
	// 	constitution := strings.TrimSpace(`
	// We the People of the United States, in Order to form a more perfect Union,
	// establish Justice, insure domestic Tranquility, provide for the common defence,
	// promote the general Welfare, and secure the Blessings of Liberty to ourselves
	// and our Posterity, do ordain and establish this Constitution for the United
	// States of America.
	// `)

	// 	got := strings.TrimSpace(`
	// :wq
	// We the People of the United States, in Order to form a more perfect Union,
	// establish Justice, insure domestic Tranquility, provide for the common defence,
	// and secure the Blessings of Liberty to ourselves
	// and our Posterity, do ordain and establish this Constitution for the United
	// States of America.
	// `)
	// 	P(diff.Diff(got, constitution))

	P(diff.Diff("old", "new"))

	// [ `run` | done: 2.495417ms ]
	// -old
	// +new
	/*

		Error loading syntax file "Packages/GoSublime/syntax/GoSublime-Go.tmLanguage":
		 Unable to open Packages/GoSublime/syntax/GoSublime-Go.tmLanguage
	*/

}
