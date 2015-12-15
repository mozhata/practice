package main

import (
	"fmt"
	"reflect"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	// basic()
	mass()
}

func basic() {
	x := 8.6
	y := float32(3.2)
	word := "China"

	P("type of x,and value: ", reflect.TypeOf(x), x)
	P("type of y,and value: ", reflect.TypeOf(y), y)

	value := reflect.ValueOf(word)
	value2 := reflect.ValueOf(y)
	text := value.String()
	text2 := value2.String()
	P(text, text2)
	P(23 << 1)
}
func mass() {
	// p := new([]int)
	// P(p)
	// v := make([]int, 10, 100)
	// P(v)

	// b := new(bool)
	// P(b)
	// P(*b)
	// i := new(int)
	// P(i)
	// P(*i)
	// s := new(string)
	// fmt.Printf("ab%vc", s)
	// P()
	// fmt.Printf("ab%sc", *s)
	// P(*s)
}
