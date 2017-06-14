package cha

import "fmt"

func TryChannel() {
	ch := make(chan int, 1024)
	for i := 0; i < 100; i++ {
		select {
		case ch <- 1:
		case ch <- 2:
		}
		i := <-ch
		fmt.Printf("value received: %v\n", i)
	}
	// for i := range ch {
	// 	fmt.Printf("value received: %v\n", i)
	// }
}
