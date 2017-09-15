/*
TODO: chapter 7
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"practice/go/gopl/chap12"
	"sort"
	"strings"
	"time"

	"github.com/golang/glog"

	"golang.org/x/net/html"
)

// client at practice/go/tools/

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	// basicFib()
	// baicNetConn()
	// echoConn()
	// TryToposort()
	// parseHtml()
	// chap12.BasicRefect()
	// chap12.TryAny()
	// chap12.TryDisplay()
	chap12.TrySMarshal()
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

// chapter5.6: 匿名函数

// 前置课程
func TryToposort() {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	fmt.Println("order is: as followed")
	for _, cls := range toposort(prereqs) {
		fmt.Println(cls)
	}
}

func toposort(m map[string][]string) []string {
	var (
		order    []string
		seen     = make(map[string]bool)
		visitAll func(items []string)
	)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	visitAll(keys)
	return order
}

// chapter5.2: 递归
func parseHtml() {
	// url := "https://golang.org"
	url := "https://baidu.com"
	glog.Infof("getting response form link %s", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	glog.Infof("status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("response not 200: %d\n", resp.StatusCode)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Printf("link:\t %s\n", link)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Printf("stach: %v\n", stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
