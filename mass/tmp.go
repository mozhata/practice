package main

import "fmt"

func main() {
	slc := []string{"a", "b"}
	fmt.Printf("slc: %v\nreverse: %v", slc, reverse(slc))
}
