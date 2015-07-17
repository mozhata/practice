package main

import (
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,
		"hello, this is a very simple website supported by Golang")
}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
