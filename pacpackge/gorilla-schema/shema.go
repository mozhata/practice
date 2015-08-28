package main

import (
	"fmt"

	"github.com/gorilla/schema"
)

var P func(...interface{}) (int, error) = fmt.Println

type Person struct {
	Name  string
	Phone string
}

func main() {
	val := map[string][]string{
		"Name":  {"John"},
		"Phone": {"999"},
	}
	person := new(Person)
	decoder := schema.NewDecoder()
	decoder.Decode(person, val)
	P(person)
}
