package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// netCatV1()
	netCatV2()
}

/*
8.3. 示例: 并发的Echo服务
*/

func netCatV2() {
	conn, err := net.Dial("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected from :8181\n")
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

/*
chapter8.2: 并发的Clock服务
*/
func netCatV1() {
	conn, err := net.Dial("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected from :8181\n")
	defer conn.Close()
	mustCopy(os.Stdout, conn)
	log.Println("exit")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
