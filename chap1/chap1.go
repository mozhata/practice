package main

import (
	"fmt"
	"image/color"
)

type ColorPoint struct {
	color.Color
	x, y int
}

const (
	c0 = iota
	c1
	c2 = 8
	c3 = iota
	c4
)

func main() {
	// fmt.Println(c0, c1, c2, c3, c4)
	// s := "abc"
	// fmt.Println(&s)

	// s, y := "hello'", 20
	// fmt.Println(&s, y)
	// {
	// 	s, z := 1000, 30
	// 	fmt.Println(&s, z)
	// }
	// str := "helo, 师姐"
	// for i, ch := range str {
	// 	fmt.Println(i, string(ch))
	// }
	slice := make([]int, 5, 10)
	for _, ch := range slice {
		fmt.Println(ch)
	}
}

// 0110	6
// 1011	116 & 11
