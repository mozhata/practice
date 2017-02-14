package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// basicFib()
	baicNetConn()
}

/*
chapter8.2: 并发的Clock服务
*/
func baicNetConn() {
	listener, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on :8181")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// handleConn
		handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Printf("write to conn failed, err: %v", err)
			return
		}
		time.Sleep(time.Second)
	}
}

/*
chapter8.1: Goroutines && Channels
*/
func basicFib() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
