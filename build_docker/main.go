package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	loger := log.New(os.Stderr, "", log.Lshortfile)
	begin := time.Now()
	loger.Println("program begins at: ", begin.Format(layout))

	// go func() {
	// 	c := time.Tick(time.Second * 3)
	// 	for cur := range c {
	// 		// loger.Printf("program update at %s, have ran %d Second (begins at %s)",
	// 		// cur.Format(layout), int(cur.Sub(begin).Seconds()), begin.Format(layout))
	// 		resp, err := http.Get("http://www.baidu.com")
	// 		loger.Printf("at %s, the request baidu err is %v, response code is %d", cur.Format(layout), err, resp.StatusCode)
	// 	}

	// }()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, now is %s\n", time.Now().Format(layout))
	})
	log.Fatal(http.ListenAndServe(":8091", mux))

}
