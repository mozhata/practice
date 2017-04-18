package tryjson

import (
	"encoding/json"
	"fmt"
)

// Empty omitempty 会去掉零值
func Empty() {
	type s struct {
		FieldOne   int    `json:"one,omitempty"`
		FieldTwo   string `json:"two,omitempty"`
		FieldThree int    `json:"three"`
	}
	s1 := s{FieldTwo: ""}
	b1, _ := json.Marshal(s1)
	fmt.Println("marshaled: ", string(b1))
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
