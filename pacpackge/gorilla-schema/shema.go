package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

var P func(...interface{}) (int, error) = fmt.Println

type Person struct {
	Name  string
	Phone string
}

func main() {
	trUrlValue()
}
func tryDecoder() {
	val := map[string][]string{
		"Name":  {"John"},
		"Phone": {"999"},
	}
	person := new(Person)
	decoder := schema.NewDecoder()
	decoder.Decode(person, val)
	P(person)
}
func trUrlValue() {
	// type Values map[string][]string
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "jess")
	v.Add("friend", "yangyang")
	v.Add("friend", "zye")

	P(v.Get("name"))
	P(v.Get("friend"))
	P(v["friend"])
	P(v)
}
