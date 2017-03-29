package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(fileName string) *os.File {
	fmt.Println("creating..")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(fileHandler *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(fileHandler, "data1~~")
	fmt.Fprintln(fileHandler, "data2~~")
}

func closeFile(fileHandler *os.File) {
	fmt.Println("closing..")
	fileHandler.Close()
}
