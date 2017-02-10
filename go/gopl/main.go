package main

import "time"

func main() {

}

/*
chapter8: Goroutines && Channels
*/
func basicFib() {
	go spinner(100 * time.Millisecond)
	const n = 45

}
