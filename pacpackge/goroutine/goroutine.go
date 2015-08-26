package main

import "fmt"

var P func(...interface{}) (int, error) = fmt.Println

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
func Count(ch chan int) {
	P("counting...")
	ch <- 1
}
