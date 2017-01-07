package main

import (
	"log"
	"os"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05.999999999"
	loger := log.New(os.Stderr, "", log.Lshortfile)
	begin := time.Now()
	loger.Println("program begins at: ", begin.Format(layout))

	c := time.Tick(time.Second * 3)
	for cur := range c {
		// cur := cur
		loger.Printf("program update at %s, have ran %d Second (begins at %s)",
			cur.Format(layout), int(cur.Sub(begin).Seconds()), begin.Format(layout))
	}
}
