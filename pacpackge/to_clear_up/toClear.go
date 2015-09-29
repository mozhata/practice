package main

import (
	"fmt"

	"bitbucket.org/applysquare/applysquare-go/pkg/common"
)

type A struct {
	Name  string `json:"name"`
	name2 string `json:"name2"`
	Name3 string
}
type B struct {
	Name3 string `json:"name3"`
	Aa    A      `josn:"a"`
	name4 string `json:"name4"` // 小写的话json中不显示
}
type C struct {
	A
	AA A `json:"aa"`
}
type M map[string]interface{}

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	var vb B
	var vc C
	var va A
	va.Name = "123"
	va.name2 = "234"
	va.Name3 = "name3"

	vb.Name3 = "vb123"
	vb.name4 = "bv345"
	vb.Aa = va

	vc.A = va
	vc.AA = va
	vc.Name = "name"
	// vc.name2 = "name2"

	jva, _ := common.CleanJSON(M{"va": va}).MarshalJSON()
	P(string(jva))
	jvb, _ := common.CleanJSON(M{"vb": vb}).MarshalJSON()
	P(string(jvb))
	jvc, _ := common.CleanJSON(M{"vc": vc}).MarshalJSON()
	P(string(jvc))

}
