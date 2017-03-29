package main

import (
	"bufio"
	"flag"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/golang/glog"
)

var (
	googleAPI = "https://www.google.co.jp/search"
	source    = "tes.txt"
	// source    = "site.txt"
)

func main() {
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()

	glog.Infoln("begin...")

	ReadLine(source, useGet)

	glog.Infoln("end")
}

type csv struct {
	*os.File
}

// func (c *csv)write

func useGet(query string) {
	uri := googleAPI + "?q=" + url.QueryEscape(query)
	glog.Infoln("uri: ", uri)
	resp, err := http.Get(uri)
	glog.Infoln("err: ", err)
	glog.Infof("resp: %#v", resp)
}

// func writeCSV() error {
// 	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
