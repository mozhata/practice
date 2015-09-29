package main

import (
	"fmt"
	"reflect"
)

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	basic()
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
}
