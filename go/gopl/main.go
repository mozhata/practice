package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// client at practice/go/tools/

func main() {
	// basicFib()
	// baicNetConn()
	// echoConn()
	pipeLine()
}

/*
chapter8.4 Channels
*/
func pipeLine() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
func counter(out chan<- int) {
	for x := 0; x < 20; x++ {
		out <- x
		time.Sleep(time.Millisecond * 300)
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
	fmt.Println("finished")
}

/*
chapter8.3. 示例: 并发的Echo服务
*/
func echoConn() {
	listener, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("echoConn listening on :8181")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// handleEchoConn
		go handleEchoConn(conn)
	}
}

func handleEchoConn(c net.Conn) {
	fmt.Fprintln(c, "handling conn..")
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	log.Println("handling conn...")
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
