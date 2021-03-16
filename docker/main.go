package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	// flag.IntVar("num", "num count", "")
	num = flag.Int("num", 20, "num of count")
)

func main() {
	flag.Parse()

	for i := 0; i < *num; i++ {
		fmt.Printf("num: %d\n", i)
		time.Sleep(time.Second)
	}
}
