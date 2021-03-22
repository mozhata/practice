package tryjson

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	FieldOne   int     `json:"one,omitempty"`
	FieldTwo   string  `json:"two,omitempty"`
	FieldThree *string `json:"three,omitempty"`
	FieldFour  string  `json:"four"`
}

// Empty omitempty 会去掉零值
func Empty() {
	type foo struct {
		Int   int     `json:"int,omitempty"`
		Pint  *int    `json:"pint,omitempty"`
		Str   string  `json:"str,omitempty"`
		Pstr  *string `json:"pstr,omitempty"`
		Bool  bool    `json:"bool,omitempty"`
		Pbool *bool   `json:"pbool,omitempty"`
	}
	var (
		zeroInt  = 0
		zeroStr  = ""
		zeroBool = false
	)
	f := foo{
		Int:   zeroInt,
		Pint:  &zeroInt,
		Str:   zeroStr,
		Pstr:  &zeroStr,
		Bool:  zeroBool,
		Pbool: &zeroBool,
	}

	b, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	jsoned := string(b)
	var fu foo
	err = json.Unmarshal([]byte(jsoned), &fu)
	fmt.Printf("instance foo:\n\t%#v\nmarshaled to json:\n\t%s\nand unmarshaled back:\n\t%#v\n",
		f, jsoned, fu)
}

func UmarshalJSON() {

	source := `{"a": "abc", "b": "bcd"}`
	printOutput := func(unitName, source string, err error, dest interface{}) {
		fmt.Printf("umarshal %s: source: %s\terr: %v\tdest: %#v\n", unitName, source, err, dest)
	}

	// struct 是source的子集
	type unitA struct {
		A string `json:"a"`
	}
	var a unitA
	err := json.Unmarshal([]byte(source), &a)
	printOutput("unitA", source, err, a)

	// struct 是source的超集
	type unitB struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	var b unitB
	err = json.Unmarshal([]byte(source), &b)
	printOutput("unitB", source, err, b)

	// struct与source有交集
	type unitC struct {
		A string `json:"a"`
		C string `json:"c"`
	}
	var c unitC
	err = json.Unmarshal([]byte(source), &c)
	printOutput("unitC", source, err, c)

}

func MarshalIntMap() {
	fmt.Printf("try marshal map[int]string:\n%s\n", MarshalJSONOrDie(map[int]string{1: "abc"}))
}

func MarshalJSONOrDie(val interface{}) []byte {
	bs, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return bs
}
