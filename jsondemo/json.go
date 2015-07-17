package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

var book Book = Book{
	"Golang programming",
	[]string{"Xu", "Lv", "Song", "Han", "Yuan", "Pan", "XuDaoli"},
	"ituring.com.cn",
	true,
	9.9,
}

func main() {
	b, err := json.Marshal(book)
	fmt.Println("err: ", err, "\nbook: \n")
	fmt.Println(string(b))
	/*	var book2 Book
		err2 := json.Unmarshal(b, &book2)
		fmt.Println("err2: ", err2, "\nbook2: \n")
		fmt.Println(book2)
	*/
	fmt.Println("\n\nnext is a sequences of test...\n\n")
	var r interface{}
	err3 := json.Unmarshal(b, &r)
	fmt.Println(err3)
	gobook2, ok := r.(map[string]interface{})
	if ok {
		for k, v := range gobook2 {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			/*case []interface{}:
			fmt.Println(k, "is an arry: ")
			for i, iv := range v2 {
				fmt.Println(i, iv)
			}*/
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}
