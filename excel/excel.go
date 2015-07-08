package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	var mySlice [][][]string
	var value string
	path := "/home/jason/go/src/fortest/testfile.xlsx"
	mySlice, _ = xlsx.FileToSlice(path)
	// value = mySlice[0][0][0]
	// fmt.Println(value, "\n there is some testings: \n")
	row1 := mySlice[0][0]
	row2 := mySlice[0][1]
	row3 := mySlice[0][2]
	fmt.Println("sheet: ", len(mySlice), "rows: ", len(mySlice[0]))
	for _, row := range mySlice[0] {
		fmt.Println("colNum-of-this-row: ", len(row))
	}
	fmt.Println("length:\t", len(row1), len(row2), len(row3), len(mySlice))
	fmt.Println("row1:\n")
	for _, val := range row1 {
		fmt.Print(val, "\t")
	}
	fmt.Print("\n")
	fmt.Println("row2:\n")
	for _, val := range row2 {
		fmt.Print(val, "\t")
	}
	fmt.Print("\n")
	fmt.Println("row3:\n")
	for _, val := range row3 {
		fmt.Print(val, "\t")
	}
	fmt.Print("\n")

}
