package main

import (
	"practice/go/encrypt/symmetric"
	"practice/go/util"
)

const (
	key = "1234567890123456"
)

func main() {
	content := "this is content"
	result, err := symmetric.Encrypt(key, content)
	util.Debug("result: %s\nerr: %v", result, err)
}
